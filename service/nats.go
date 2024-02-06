package service

import (
	"context"
	"fmt"
	"insellisense-mqtt-transport-service/config"

	"github.com/nats-io/nats.go"
)

type Nats struct {
	Nc *nats.Conn
}

func NewNatsService(ctx context.Context, conf config.Nats) (*Nats, error) {
	natsService := new(Nats)

	err := natsService.connectToNatsServer(ctx, conf.Url, conf.Username, conf.Password)
	if err != nil {
		return nil, err
	}

	return natsService, nil
}

func (n *Nats) connectToNatsServer(ctx context.Context, url string, user string, pass string) error {
	nc, err := nats.Connect(url, nats.UserInfo(user, pass))
	if err != nil {
		return fmt.Errorf("error when connect to NATS Server, url: %s, user: %s. error: %w", url, user, err)
	}

	n.Nc = nc
	return nil
}
