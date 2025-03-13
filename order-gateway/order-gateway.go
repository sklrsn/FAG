package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	otellib "github.com/sklrsn/FAG/lib/otel"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
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
	otelShutdownFunc, err := otellib.SetupOTel(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer otelShutdownFunc(context.Background())

	router := mux.NewRouter()

	router.HandleFunc("/order-gateway/buy", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "/order-gateway/buy")
		defer span.End()

		var po PurchaseOrder
		if err := json.NewDecoder(r.Body).Decode(&po); err != nil {
			http.Error(w, "invalid request data", http.StatusBadRequest)
			return
		}

		// SAGA Pattern
		sagaOrchestrator := new(SagaOrchestrator)

		orderID, err := sagaOrchestrator.Orchestrate(ctx, po)
		if err != nil {
			http.Error(w, "failed to process order", http.StatusBadRequest)
			return
		}

		data, err := json.Marshal(Order{
			OrderID: orderID,
		})
		if err != nil {
			http.Error(w, "failed to marshall order data", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(data); err != nil {
			http.Error(w, "partial response", http.StatusBadRequest)
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
