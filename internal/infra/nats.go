package infra

import (
	"github.com/nats-io/nats.go"
)

func NewNats(config *FlowConfig) (*nats.Conn, error) {
	natsUrl := config.NatsUrl
	if natsUrl == "" {
		natsUrl = nats.DefaultURL
	}

	nc, err := nats.Connect(natsUrl)

	return nc, err
}
