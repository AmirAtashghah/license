package userservice

import (
	"github.com/go-playground/validator/v10"
	"server/entity"
	"server/logger"
	"server/pkg/param"
)

const group = "userservice"

type UserRepo interface {
	GetByUsername(username string) (entity.User, error)
}

type Service struct {
	repo UserRepo
}

func New(repo UserRepo) *Service {
	return &Service{repo: repo}
}

func (s Service) ValidateLoginRequest(lr *param.LoginRequest) error {

	validate := validator.New()

	if err := validate.Struct(lr); err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error())

		return err
	}

	return nil
}

func (s Service) GetUserByUsername(username string) (entity.User, error) {

	// todo impolent it

	return entity.User{}, nil
}
