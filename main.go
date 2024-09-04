package main

import (
	"context"
	"log"
	"otelprac2/logging"
	"otelprac2/metrics"
	"otelprac2/tracing"
	"github.com/rpradeepkumar7/otelprac2/logging"
	"github.com/rpradeepkumar7/otelprac2/metrics"
	"github.com/rpradeepkumar7/otelprac2/tracing"
)

func main() {
	ctx := context.Background()

	// Initialize tracing, metrics, and logging
	tracing.InitTracerProvider(ctx)
	metrics.InitMeterProvider(ctx)
	logging.InitLogger(ctx)

	// Example usage of tracing, metrics, and logging
	tracing.StartSpan(ctx, "example-span")
	metrics.RecordMetrics(ctx)
	logging.LogEvent(ctx)

	log.Println("OpenTelemetry setup completed with the new JSON data model!")

	logging.LogSomething()
    	metrics.RecordMetrics()
    	tracing.StartTrace()
}
