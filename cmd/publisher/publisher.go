package main

import (
	"time"

	"github.com/kovalyov-valentin/orders-service/pkg/lib/random"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {

	sc, err := stan.Connect(
		"test-cluster",
		"publisher",
		stan.NatsURL("0.0.0.0:4222"),
	)
	if err != nil {
		logrus.Fatal("error connect to nats-streaming", err)
		return
	}
	defer sc.Close()

	t := time.NewTicker(2 * time.Second)
	for range t.C {
		logrus.Println("send order to orders")
		
		if err = sc.Publish("order", random.RandOrder()); err != nil {
			logrus.Fatal(err)
		}
	}
}

