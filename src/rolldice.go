package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const name = "github.com/ZiadMansour/bastet/examples/dice"

var (
	tracer  = otel.Tracer(name)
	meter   = otel.Meter(name)
	logger  = otelslog.NewLogger(name)
	rollCnt metric.Int64Counter
)

func init() {
	var err error
	rollCnt, err = meter.Int64Counter("dice.rolls",
		metric.WithDescription("The number of rolls by roll value"),
		metric.WithUnit("{roll}"))
	if err != nil {
		panic(err)
	}
}

func rolldice(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "roll_dice")
	defer span.End()

	// Add an event to indicate the dice roll has started
	span.AddEvent("Dice rolling started")

	// Roll the dice four times and collect the results
	results := make([]int, 4)
	for i := 0; i < 4; i++ {
		results[i] = rollSingleDice(ctx, i+1)
	}

	// Add an event to indicate all rolls are complete
	span.AddEvent(
		"Dice rolling completed",
		trace.WithAttributes(attribute.IntSlice("results", results)),
	)

	// Generate a response message
	response := fmt.Sprintf("Dice rolls: %v\n", results)
	if _, err := io.WriteString(w, response); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to write response")
		log.Printf("Write failed: %v\n", err)
		return
	}

	span.SetStatus(codes.Ok, "Roll dice handler completed successfully")
}

func rollSingleDice(ctx context.Context, rollNumber int) int {
	// Start a child span for this roll
	ctx, span := tracer.Start(ctx, fmt.Sprintf("roll_%d", rollNumber))
	defer span.End()

	// Add an event to indicate the roll has started
	span.AddEvent("Roll started", trace.WithAttributes(attribute.Int("roll.number", rollNumber)))

	// Simulate rolling a dice
	roll := 1 + rand.Intn(6)

	// Add attributes and metrics for this roll
	rollValueAttr := attribute.Int("roll.value", roll)
	span.SetAttributes(
		rollValueAttr,
		attribute.Int("roll.number", rollNumber),
	)
	rollCnt.Add(ctx, 1, metric.WithAttributes(rollValueAttr))

	// Log the roll result
	logger.InfoContext(ctx, "Dice rolled", "roll_number", rollNumber, "result", roll)
	span.AddEvent("Roll completed", trace.WithAttributes(rollValueAttr))

	// Simulate an exception for certain conditions (example: roll of 1)
	if roll == 1 {
		err := fmt.Errorf("unlucky roll: %d", roll)
		span.RecordError(err)
		span.SetStatus(codes.Error, "Critical roll error")
		logger.ErrorContext(ctx, "Critical roll error", "roll_number", rollNumber, "error", err)
	}

	return roll
}
