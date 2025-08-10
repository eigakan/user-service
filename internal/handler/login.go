package handler

import (
	"encoding/json"
	"fmt"
	"time"

	dto "github.com/eigakan/nats-shared/dto/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nats-io/nats.go"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandlers) makeJwtClaim() *jwt.MapClaims {
	return &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(h.jwtConf.ExpHours)).Unix(),
	}
}

func (h *UserHandlers) Login(msg *nats.Msg) {
	var reqDto dto.LoginRequestDTO

	if err := json.Unmarshal(msg.Data, &reqDto); err != nil {
		h.nc.RespondErr(msg, "Invalid payload")
		return
	}

	user, err := h.r.GetUserByLogin(reqDto.Login)
	if err != nil {
		fmt.Print(err)
		h.nc.RespondErr(msg, "Wrong password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqDto.Password))
	if err != nil {
		h.nc.RespondErr(msg, "Wrong password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, h.makeJwtClaim())
	tokenString, err := token.SignedString([]byte(h.jwtConf.Secret))

	fmt.Println(tokenString)

	if data, err := json.Marshal(dto.LoginResponseDTO{Token: tokenString}); err == nil {
		h.nc.Respond(msg, data)
		return
	}
}
