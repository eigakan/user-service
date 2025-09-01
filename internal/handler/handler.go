package handler

import (
	nats_client "github.com/eigakan/nats-shared/client"
	"github.com/eigakan/nats-shared/topics"
	"github.com/eigakan/user-service/config"
	"github.com/eigakan/user-service/internal/repository"
)

type UserHandlers struct {
	nc      *nats_client.Client
	r       *repository.UserRepository
	jwtConf *config.JwtConfig
}

func NewUserHandlers(nc *nats_client.Client, r *repository.UserRepository, jwtConf *config.JwtConfig) *UserHandlers {
	return &UserHandlers{nc: nc, r: r, jwtConf: jwtConf}
}

func (h *UserHandlers) RegisterHandlers() {
	h.nc.Subscribe(topics.UserGet, h.GetUser)
	h.nc.Subscribe(topics.UserCreate, h.CreateUser)
	h.nc.Subscribe(topics.UserGetByPassword, h.GetUserByPassword)
}
