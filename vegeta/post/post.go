package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kovalyov-valentin/orders-service/internal/models"
	"github.com/kovalyov-valentin/orders-service/utils"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	attackName := "POST ORDER"
	rate := vegeta.Rate{Freq: 4000, Per: time.Second}
	duration := 10 * time.Second

	targeter := getNewTargeter()
	attacker := vegeta.NewAttacker()
	
	var metrics vegeta.Metrics
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	var successCount int64 = 0
	for res := range attacker.Attack(targeter, rate, duration, attackName) {
		metrics.Add(res)
		if res.Code == http.StatusOK {
			successCount++
		}
	}
	utils.PrintMetrics(attackName, &metrics, successCount, rate.String())
	defer metrics.Close()
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
    fmt.Printf("50th percentile: %s\n", metrics.Latencies.P50)
    fmt.Printf("Mean: %s\n", metrics.Latencies.Mean)
    fmt.Printf("Max: %s\n", metrics.Latencies.Max)
    fmt.Printf("Total Duration: %s\n", metrics.Duration)
}

func getNewTargeter() vegeta.Targeter {
	return func(tg *vegeta.Target) error {
		if tg == nil {
			return vegeta.ErrNilTarget
		}
		socket := "http://localhost:8080"
		targetAPI := "orders"
		URL := fmt.Sprintf("%s/%s", socket, targetAPI)

		tg.Method = http.MethodPost
		tg.URL = URL
		tg.Header = http.Header{
			"Content-Type": {"application/json"},
		}
		tg.Body = models.GetRandomOrderEncodedJson()
		return nil
	}
}
