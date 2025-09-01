package handler

import (
	"encoding/json"

	dto "github.com/eigakan/nats-shared/dto/user"
	nats_model "github.com/eigakan/nats-shared/model"
	"github.com/eigakan/user-service/internal/model"
	"github.com/nats-io/nats.go"
)

func (h *UserHandlers) GetUser(msg *nats.Msg) {
	var reqDto dto.GetUserRequestDTO

	if err := json.Unmarshal(msg.Data, &reqDto); err != nil {
		h.nc.RespondErr(msg, "Invalid payload")
		return
	}

	var user *model.User

	user, _ = h.r.GetUserById(reqDto.UserID)

	if user == nil {
		h.nc.RespondErr(msg, "No user found by provded id or login")
		return
	}

	h.nc.Respond(msg, dto.GetUserResponseDTO{
		User: nats_model.User{
			ID:        user.ID,
			Email:     user.Email,
			Login:     user.Login,
			Logo:      user.Logo,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}
