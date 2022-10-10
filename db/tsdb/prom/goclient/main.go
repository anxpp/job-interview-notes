package main

import (
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
}

func recordMetrics() {
	//每2秒，计数器增加1。
	go func() {
		for {
			opsProcessed.Inc()
			cpuTemp.Set(rand.Float64() * 100)
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "myapp_processed_ops_total",
		Help:        "The total number of processed events",
		ConstLabels: map[string]string{"key": "value"},
	})
)

func main() {
	recordMetrics()
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8888", nil))
}
