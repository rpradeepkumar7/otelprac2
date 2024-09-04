package metrics

import (
	"context"
	"log"
	"time"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/metric"
)

func InitMeterProvider(ctx context.Context) {
	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", "web-backend"),
			attribute.String("host.name", "web-server-1"),
			attribute.String("host.ip", "192.168.1.1"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	provider := metric.NewMeterProvider(
		metric.WithResource(res),
	)
	otel.SetMeterProvider(provider)
}

func RecordMetrics(ctx context.Context) {
	meter := otel.Meter("example-meter")
	latency := meter.NewFloat64Histogram("network.latency")

	latency.Record(ctx, 100.5, attribute.String("host.name", "web-server-1"))
	time.Sleep(2 * time.Second)
}
