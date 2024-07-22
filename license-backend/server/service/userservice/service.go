package userservice

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"server/entity"
	"server/pkg/param"
)

type UserRepo interface {
	GetUserByUsername(username string) (entity.User, error)
}

type Service struct {
	repo UserRepo
}

func (s Service) ValidateLoginRequest(lr *param.LoginRequest) error {

	validate := validator.New()

	if err := validate.Struct(lr); err != nil {
		return fiber.NewError(400, err.Error())
	}

	return nil
}

func (s Service) GetUserByUsername(username string) (entity.User, error) { return entity.User{}, nil }
