package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kovalyov-valentin/orders-service/internal"
	"github.com/kovalyov-valentin/orders-service/internal/repository/cache"
	config "github.com/kovalyov-valentin/orders-service/internal/config"
	"github.com/kovalyov-valentin/orders-service/internal/models"
	"github.com/kovalyov-valentin/orders-service/internal/handler"
	"github.com/kovalyov-valentin/orders-service/internal/repository"
	"github.com/kovalyov-valentin/orders-service/internal/service"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg := config.MustLoad()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CtxTimeout)
	defer cancel()

	db, err := repository.NewPostgresDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to connect postgres db: %s", err.Error())
	}
	defer db.Close()

	cache := cache.NewCache()
	repos := repository.NewRepository(db)
	services := service.NewService(repos, cache)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)

	serverErrors := make(chan error, 1)
	go func() {
		logrus.Printf("Start listen http service on %s at %s\n", cfg.Address, time.Now().Format(time.DateTime))
		err := srv.Run(cfg.HTTPServer, handlers.InitRoutes())
		if err != nil {
			logrus.Printf("shutting down the server: %s\n", cfg.Address)
		}
		serverErrors <- err
	}()

	nc, err := stan.Connect(
		"test-cluster",
		"subscriber",
		stan.NatsURL("0.0.0.0:4222"),
	)
	if err != nil {
		logrus.Fatal("Can't connect to nats streaming server with error: " + err.Error())
		return 
	}
	

	defer nc.Close()

	_, err = nc.Subscribe("order", func(m *stan.Msg) {
		fmt.Print(string(m.Data))
		var order models.Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			logrus.Println("consumer: error unmarshal json")
			return
		}
		if err = (*service.OrderService).Create(services.ServiceOrder, ctx, order.OrderUID, order); err != nil {
			return
		}
	})
	if err != nil {
		logrus.Println("error subscribe")
	}

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err := <-serverErrors:
		logrus.Printf("error starting server: %v\n", err)
	case <-osSignal:
		logrus.Println("start shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), cfg.CtxTimeout)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logrus.Printf("graceful shutdown error: %v\n", err)
			os.Exit(1)
		}
	}
}
