package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/pkg/jwt"
	"server/pkg/param"
)

// todo add log

func (h Handler) GetTokenFromCookie(ctx *fiber.Ctx) error {

	authHeader := ctx.Cookies("token", "")

	if authHeader == "" {

		return ctx.Status(fiber.ErrForbidden.Code).JSON(fiber.Map{"error": "forbidden"})
	}

	userID, role, err := getTokenInfo(authHeader, h.jwtSvc)
	if err != nil {
		return ctx.Status(fiber.ErrForbidden.Code).JSON(fiber.Map{"error": "forbidden"})
	}

	ctx.Locals("userID", userID)
	ctx.Locals("role", role)

	return ctx.Next()
}

func getTokenInfo(token string, j *jwt.Service) (string, string, error) {
	c, err := j.VerifyToken(token)
	if err != nil {

		return "", "", err
	}

	m := j.GetClaimsValue(c)
	_, ok := m["userID"]
	if !ok {
		return "", "", fmt.Errorf("can not get token claims")
	}

	_, ok = m["role"]
	if !ok {
		return "", "", fmt.Errorf("can not get token claims")
	}

	return m["userID"], m["role"], nil
}

func (h Handler) ValidateToken(ctx *fiber.Ctx) error {

	req := new(param.ValidateTokenRequest)

	if err := json.Unmarshal(ctx.Body(), req); err != nil {

		return ctx.Status(fiber.ErrForbidden.Code).JSON(fiber.Map{"error": "forbidden"})
	}

	_, err := h.jwtSvc.VerifyToken(req.Token)
	if err != nil {

		return ctx.Status(fiber.ErrForbidden.Code).JSON(fiber.Map{"error": "forbidden"})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) SuperAdminCheck(ctx *fiber.Ctx) error {
	// get role from token
	userRole := ctx.Locals("role")

	if userRole == "superAdmin" {
		return ctx.Next()
	}

	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
}
