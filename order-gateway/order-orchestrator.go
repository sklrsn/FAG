package main

import (
	"fmt"
)

type SagaOrchestrator struct {
}

func (so *SagaOrchestrator) Process(order PurchaseOrder) (string, error) {
	return "", fmt.Errorf("not implemented")
}
