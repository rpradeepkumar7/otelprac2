package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/attribute"
	"log"
)

func InitTracerProvider(ctx context.Context) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", "web-backend"),
			attribute.String("host.name", "web-server-1"),
			attribute.String("host.ip", "192.168.1.1"),
			attribute.String("host.mac", "00:1A:2B:3C:4D:5E"),
			attribute.String("os.type", "Linux"),
			attribute.String("os.version", "Ubuntu 20.04"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)
}

func StartSpan(ctx context.Context, name string) {
	tracer := otel.Tracer("example-tracer")
	_, span := tracer.Start(ctx, name)
	defer span.End()
}
