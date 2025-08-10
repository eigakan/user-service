package handler

import (
	"encoding/json"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/nats-io/nats.go"
)

func (h *UserHandlers) GetUser(msg *nats.Msg) {
	var reqDto dto.GetUserRequestDTO

	if err := json.Unmarshal(msg.Data, &reqDto); err != nil {
		h.nc.RespondErr(msg, "Invalid payload")
		return
	}

	user, err := h.r.GetUserByLogin(reqDto.Login)
	if err != nil {
		h.nc.RespondErr(msg, "No such user")
		return
	}

	h.nc.Respond(msg, dto.GetUserResponseDTO{
		ID:        user.ID,
		Email:     user.Email,
		Login:     user.Login,
		Logo:      user.Logo,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
