package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/sklrsn/gRPC-defs/order"
	"github.com/sklrsn/gRPC-defs/payment"
	"github.com/sklrsn/gRPC-defs/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SagaOrchestrator struct {
	orderClient    order.OrderEngineClient
	paymentClient  payment.PaymentClient
	shippingClient shipping.ShippingClient
}

var (
	orderEngineUrl    = os.Getenv("ORDER_ENGINE_ADDRESS")    //"localhost:9091"
	paymentEngineUrl  = os.Getenv("PAYMENT_ENGINE_ADDRESS")  //"localhost:9092"
	shippingEngineUrl = os.Getenv("SHIPPING_ENGINE_ADDRESS") //"localhost:9093"
)

func (so *SagaOrchestrator) Init() error {
	oeCl, err := grpc.NewClient(orderEngineUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	so.orderClient = order.NewOrderEngineClient(oeCl)

	payCl, err := grpc.NewClient(paymentEngineUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	so.paymentClient = payment.NewPaymentClient(payCl)

	shipCl, err := grpc.NewClient(shippingEngineUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	so.shippingClient = shipping.NewShippingClient(shipCl)

	return nil
}

func (so *SagaOrchestrator) Orchestrate(ctx context.Context, po PurchaseOrder) (string, error) {
	userId := uuid.NewString()
	productId := uuid.NewString()
	merchantId := uuid.NewString()

	reserveResponse, err := so.orderClient.Reserve(ctx,
		&order.ReserveRequest{
			UserId:       userId,
			ProductId:    productId,
			ReserveCount: 3})
	if err != nil {
		return "", err
	}
	log.Printf("reserve response :%v", reserveResponse)

	preAuthResponse, err := so.paymentClient.PreAuthorize(ctx, &payment.PreAuthorizeRequest{
		UserId:     userId,
		MerchantId: merchantId,
		Amount:     300,
	})
	if err != nil {
		return "", err
	}
	log.Printf("preauth response :%v", preAuthResponse)

	captureResponse, err := so.paymentClient.Capture(ctx,
		&payment.CaptureRequest{
			PreAuthorizationId: preAuthResponse.PreAuthorizationId})
	if err != nil {
		return "", err
	}
	log.Printf("capture response :%v", captureResponse)

	orderId := uuid.NewString()
	dispatchResponse, err := so.shippingClient.Dispatch(ctx, &shipping.DispatchRequest{OrderId: orderId})
	if err != nil {
		return "", err
	}
	log.Printf("dispatch response :%v", dispatchResponse)

	return orderId, nil
}
