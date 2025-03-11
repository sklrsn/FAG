package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/order-gateway/buy", func(w http.ResponseWriter, r *http.Request) {
		var po PurchaseOrder
		if err := json.NewDecoder(r.Body).Decode(&po); err != nil {
			http.Error(w, "invalid request data", http.StatusBadRequest)
			return
		}

		// SAGA Pattern
		sagaOrchestrator := new(SagaOrchestrator)

		orderID, err := sagaOrchestrator.Process(po)
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

	log.Fatalf("%v", http.ListenAndServe(":8080", router))
}
