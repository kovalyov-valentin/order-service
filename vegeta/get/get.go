package main

import (
	"fmt"
	"net/http"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kovalyov-valentin/orders-service/utils"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)


func main() {
	attackName := "GET ORDER BY UID"
    rate := vegeta.Rate{Freq: 40000, Per: time.Second}
    duration := 10 * time.Second
    targeter := vegeta.NewStaticTargeter(vegeta.Target{
        Method: "GET",
        URL:    "http://localhost:8080/orders",
    })
    attacker := vegeta.NewAttacker()
 
    var metrics vegeta.Metrics
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
 
    // go func() {
    //     for range time.Tick(time.Second) {
    //         fmt.Printf("QPS: %v\n", metrics.Rate)
    //     }
    // }()
 
	var successCount int64 = 0
    for res := range attacker.Attack(targeter, rate, duration, attackName) {
        metrics.Add(res)
		if res.Code == http.StatusOK {
			successCount++
		}
    }
 
	fmt.Println(metrics.Latencies.P90, metrics.Latencies.Total)
	utils.PrintMetrics(attackName, &metrics, successCount, rate.String())
    metrics.Close()
    fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
    fmt.Printf("50th percentile: %s\n", metrics.Latencies.P50)
    fmt.Printf("Mean: %s\n", metrics.Latencies.Mean)
    fmt.Printf("Max: %s\n", metrics.Latencies.Max)
    fmt.Printf("Total Duration: %s\n", metrics.Duration)
}