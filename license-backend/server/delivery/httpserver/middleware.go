package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"server/logger"
	"server/pkg/jwt"
)

func (h Handler) GetTokenFromCookie(ctx *fiber.Ctx) error {

	authHeader := ctx.Cookies("token", "")

	if authHeader == "" {
		logger.L().WithGroup(group).Error("error", "error", "invalid token", "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "invalid token"})
	}

	userId, err := getTokenInfo(authHeader, h.jwtSvc)
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error(), "code", fiber.ErrBadRequest.Code)

		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	ctx.Set("userId", userId)

	return ctx.Next()
}

func getTokenInfo(token string, j *jwt.Service) (string, error) {
	c, err := j.VerifyToken(token)
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error())

		return "", err
	}

	m := make(map[string]string)
	m = j.GetClaimsValue(c)

	return m["userId"], nil
}
