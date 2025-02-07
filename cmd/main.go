package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	om "go.opentelemetry.io/otel/metric"
	sdkm "go.opentelemetry.io/otel/sdk/metric"
)

var meter om.Meter

func main() {
	// Setup Prometheus exporter
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("failed to initialize Prometheus exporter: %v", err)
	}

	provider := sdkm.NewMeterProvider(sdkm.WithReader(exporter))
	otel.SetMeterProvider(provider)
	meter = provider.Meter("example-app")

	// Create a sample counter
	counter, err := meter.Int64Counter("example_counter")
	if err != nil {
		log.Fatalf("failed to create counter: %v", err)
	}

	// Start HTTP server for metrics exposure
	http.Handle("/metrics", promhttp.Handler())

	// Simulate metric increments
	go func() {
		for {
			counter.Add(context.Background(), 1)
			time.Sleep(2 * time.Second)
		}
	}()

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
