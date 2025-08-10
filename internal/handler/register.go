package handler

import (
	"encoding/json"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/eigakan/user-service/internal/model"
	"github.com/nats-io/nats.go"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandlers) Register(msg *nats.Msg) {
	var reqDto dto.RegisterRequestDTO

	if err := json.Unmarshal(msg.Data, &reqDto); err != nil {
		h.nc.RespondErr(msg, "Invalid payload")
		return
	}

	pHash, err := bcrypt.GenerateFromPassword([]byte(reqDto.Password), bcrypt.DefaultCost)
	if err != nil {
		h.nc.RespondErr(msg, "Error while hashing password")
		return
	}

	user := model.User{
		Login:    reqDto.Login,
		Password: string(pHash),
		Email:    reqDto.Email,
	}

	err = h.r.Create(&user)
	if err != nil {
		h.nc.RespondErr(msg, "Error while creating user")
		return
	}

	if data, err := json.Marshal(&dto.RegisterResponseDTO{Ok: true}); err == nil {
		h.nc.Respond(msg, data)
		return
	}
}
