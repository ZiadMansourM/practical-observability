package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	serviceName      = "dice-client"
	otelCollectorURL = "localhost:4318"
)

func main() {
	// Set up OpenTelemetry
	shutdown := setupOpenTelemetry()
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shut down OpenTelemetry: %v", err)
		}
	}()

	// Create an instrumented HTTP client
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// Make 5 requests to the dice server
	for i := 0; i < 1_000; i++ {
		callDiceServer(context.Background(), &client)
		// time.Sleep(500 * time.Millisecond) // Add delay between calls to visualize traces more clearly
	}
}

func callDiceServer(ctx context.Context, client *http.Client) {
	// Start a span for the client request
	tr := otel.Tracer(serviceName)
	ctx, span := tr.Start(ctx, "CallDiceServer")
	defer span.End()

	// Make the HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:8080/rolldice/", nil)
	if err != nil {
		span.RecordError(err)
		log.Printf("Failed to create request: %v\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		span.RecordError(err)
		log.Printf("Request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		log.Printf("Failed to read response: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", body)
}

func setupOpenTelemetry() func(context.Context) error {
	// Set up an OTLP exporter to send traces to the OpenTelemetry Collector
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(otelCollectorURL),
		otlptracehttp.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("Failed to create OTLP exporter: %v", err)
	}

	// Create a TracerProvider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)

	// Set the global tracer provider and propagator
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Return a shutdown function to clean up
	return tp.Shutdown
}
