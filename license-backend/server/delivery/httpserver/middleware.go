package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/jwt"
)

func (h Handler) GetTokenFromCookie(ctx *fiber.Ctx) error {

	authHeader := ctx.Cookies("token", "")

	if authHeader == "" {
		return fiber.NewError(400, "invalid token")
	}

	userId, err := getTokenInfo(authHeader, h.jwtSvc)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Set("userId", userId)

	return ctx.Next()
}

func getTokenInfo(token string, j *jwt.Service) (string, error) {
	c, err := j.VerifyToken(token)
	if err != nil {
		return "", err
	}

	m := make(map[string]string)
	m = j.GetClaimsValue(c)

	return m["userId"], nil
}

const (
	RoleAdmin            string = "admin"
	RoleHeadofDepartment string = "headofDepartment"
	RoleMaster           string = "master"
	RoleStudent          string = "student"
)
