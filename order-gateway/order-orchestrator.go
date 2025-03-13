package main

import (
	"context"
	"fmt"
)

type SagaOrchestrator struct {
}

func (so *SagaOrchestrator) Process(ctx context.Context, order PurchaseOrder) (string, error) {
	return "", fmt.Errorf("not implemented")
}
