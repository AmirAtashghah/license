package httpserver

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"server/logger"
	"server/pkg/hash"
	"server/pkg/param"
	"time"
)

func (h Handler) Login(ctx *fiber.Ctx) error {

	lr := new(param.LoginRequest)

	if err := json.Unmarshal(ctx.Body(), lr); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if err := h.userSvc.ValidateLoginRequest(lr); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	user, err := h.userSvc.GetUserByUsername(lr.Username)
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	if !hash.CheckHash(lr.Password, user.Password) {
		logger.L().WithGroup(group).Error("error", "error", "invalid password or username", "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "invalid password or username"})
	}

	token, err := h.jwtSvc.GenerateJWTAccessToken(user.ID)
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(15 * time.Minute)

	ctx.Cookie(cookie)

	return ctx.Status(200).JSON(fiber.Map{"message": "login successfully"})
}
