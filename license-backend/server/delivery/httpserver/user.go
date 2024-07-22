package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/pkg/hash"
	"server/pkg/param"
	"time"
)

func (h Handler) Login(ctx *fiber.Ctx) error {

	lr := new(param.LoginRequest)

	if err := json.Unmarshal(ctx.Body(), lr); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if err := h.userSvc.ValidateLoginRequest(lr); err != nil {
		return fiber.NewError(400, err.Error())
	}

	user, err := h.userSvc.GetUserByUsername(lr.Username)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	if !hash.CheckHash(lr.Password, user.Password) {
		return fiber.NewError(400, fmt.Errorf("invalid password or username").Error())
	}

	token, err := h.jwtSvc.GenerateJWTAccessToken(user.ID)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(15 * time.Minute)

	ctx.Cookie(cookie)

	return ctx.Status(200).JSON(fiber.Map{"message": "login successfully"})
}
