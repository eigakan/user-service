package handler

import (
	"encoding/json"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/nats-io/nats.go"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandlers) CheckPassword(msg *nats.Msg) {
	var reqDto dto.CheckPasswordRequestDTO

	if err := json.Unmarshal(msg.Data, &reqDto); err != nil {
		h.nc.RespondErr(msg, "Invalid payload")
		return
	}

	user, err := h.r.GetUserByLogin(reqDto.Login)
	if err != nil {
		h.nc.RespondErr(msg, "No such user")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqDto.Password))
	if err != nil {
		h.nc.RespondErr(msg, "Wrong password")
		return
	}

	h.nc.Respond(msg, dto.CheckPasswordResponseDTO{Valid: true})
}
