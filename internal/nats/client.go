package nats

import (
	"fmt"

	"github.com/eigakan/user-service/config"
	"github.com/nats-io/nats.go"
)

func NewClient(config config.NatsConfig) (*nats.Conn, error) {
	return nats.Connect(fmt.Sprintf("nats://%s:%s", config.Host, config.Port))
}
