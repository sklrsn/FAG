package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func init() {}

type PurchaseOrder struct {
	Portrait string `json:"portrait_name"`
	Price    int64  `json:"portrait_price"`
	Seller   string `json:"portrait_seller,omitempty"`
}

type Order struct {
	OrderID string `json:"order_id"`
}

const ServiceName = "order-gateway"

var (
	tracer = otel.Tracer(ServiceName)
)

func main() {
	if err := SetupOTel(context.Background()); err != nil {
		log.Fatalf("%v", err)
	}

	router := mux.NewRouter()
	// SAGA Pattern
	sagaOrchestrator := new(SagaOrchestrator)
	if err := sagaOrchestrator.Init(); err != nil {
		panic(err)
	}

	router.HandleFunc("/order-gateway/buy", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "/order-gateway/buy")
		defer span.End()

		var po PurchaseOrder
		if err := json.NewDecoder(r.Body).Decode(&po); err != nil {
			http.Error(w, fmt.Sprintf("invalid request data:%v", err), http.StatusBadRequest)
			return
		}

		orderID, err := sagaOrchestrator.Orchestrate(ctx, po)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to process order:%v", err), http.StatusBadRequest)
			return
		}

		data, err := json.Marshal(Order{
			OrderID: orderID,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshall order data:%v", err), http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(data); err != nil {
			http.Error(w, fmt.Sprintf("partial response:%v", err), http.StatusBadRequest)
			return
		}
	}).Methods(
		http.MethodPost,
	)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", 8080),
		Handler: otelhttp.NewHandler(router, "/"),
	}
	server.RegisterOnShutdown(func() {
		log.Println("order-gateway is shutting down")
	})

	log.Fatalf("%v", server.ListenAndServe())
}

var (
	otelEndpoint = os.Getenv("OTEL_ENDPOINT")
)

func SetupOTel(ctx context.Context) error {
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otelEndpoint), //otlptracegrpc.WithEndpoint("localhost:4317")
	)

	if err != nil {
		return fmt.Errorf("failed to create OTLP exporter: %w", err)
	}
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("order-gateway-service"),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	return nil
}
