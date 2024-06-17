package nats

import (
	"context"
	"log/slog"
	"order/internal/service"
	"order/pkg/logger"

	"github.com/nats-io/stan.go"
)

type NatsSub struct {
	log     *slog.Logger
	service service.Order
}

func NewNatsSub(log *slog.Logger, service service.Order) *NatsSub {
	return &NatsSub{
		log:     log,
		service: service,
	}
}

func (n *NatsSub) MessageHandler(m *stan.Msg) {
	n.log.Info("Received message:", slog.String("data", string(m.Data)))

	if err := n.service.CreateOrder(context.Background(), m.Data); err != nil {
		n.log.Error("Failed to create order:", logger.Err(err))
	}
}

func (n *NatsSub) NatsConnect(url string) error {
	sc, err := stan.Connect("test-cluster", "subscriber-client", stan.NatsURL(url))
	if err != nil {
		n.log.Error("Failed to connect to NATS:", slog.String("url", url), logger.Err(err))
		return err
	}
	defer func() {
		err = sc.Close()
		if err != nil {
			n.log.Error("Failed to close NATS:", slog.String("url", url), logger.Err(err))
		}
	}()

	sub, err := sc.Subscribe("order_create", n.MessageHandler)
	if err != nil {
		n.log.Error("Failed to subscribe to NATS:", slog.String("url", url), logger.Err(err))
		return err
	}
	defer func() {
		err = sub.Unsubscribe()
		if err != nil {
			n.log.Error("Failed to unsubscribe from NATS:", slog.String("url", url), logger.Err(err))
		}
	}()

	select {}

}
