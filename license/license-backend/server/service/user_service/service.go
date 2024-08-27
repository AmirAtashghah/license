package user_service

import (
	"log"
	"server/entity"
	"server/pkg/hash"
	"server/pkg/param"
	"time"
)

type UserRepo interface {
	Insert(user *entity.User) error
	Update(user *entity.User) error
	Delete(id int) error
	GetByUsername(username string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	GetByID(id int) (*entity.User, error)
}

type Service struct {
	repo UserRepo
}

func New(repo UserRepo) *Service {
	return &Service{repo: repo}
}

func (p Service) AddUser(req *param.CreateUserRequest) error {

	hashPass, err := hash.Hash(req.Password)
	if err != nil {

		return err
	}

	user := &entity.User{
		Name:      req.Name,
		Username:  req.Username,
		Password:  hashPass,
		Role:      req.Role,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: -1,
	}
	if err := p.repo.Insert(user); err != nil {

		return err
	}

	return nil
}

func (p Service) UpdateUser(req *param.UpdateUserRequest) error {

	log.Println(req)

	var hashPass string
	var err error

	if req.Password != "" {
		hashPass, err = hash.Hash(req.Password)
		if err != nil {

			return err
		}

		user := &entity.User{
			ID:        req.ID,
			Name:      req.Name,
			Username:  req.Username,
			Password:  hashPass,
			Role:      req.Role,
			CreatedAt: req.CreatedAt,
			UpdatedAt: time.Now().Unix(),
		}

		if err := p.repo.Update(user); err != nil {

			return err
		}

	} else {

		oldUser, err := p.repo.GetByID(req.ID)
		if err != nil {
			return err
		}

		user := &entity.User{
			ID:        req.ID,
			Name:      req.Name,
			Username:  req.Username,
			Password:  oldUser.Password,
			Role:      req.Role,
			CreatedAt: req.CreatedAt,
			UpdatedAt: time.Now().Unix(),
		}

		if err := p.repo.Update(user); err != nil {

			return err
		}

	}

	return nil
}

func (p Service) DeleteUser(req *param.DeleteUserRequest) error {

	if err := p.repo.Delete(req.ID); err != nil {

		return err
	}

	return nil
}

func (p Service) GetUserByUsername(req *param.GetUserRequest) (*entity.User, error) {

	user, err := p.repo.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p Service) GetAll() ([]*entity.User, error) {

	users, err := p.repo.GetAll()
	if err != nil {

		return nil, err
	}

	return users, nil
}

func (p Service) CheckExist(username string) (bool, error) {

	user, err := p.repo.GetByUsername(username)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, nil
	}

	return true, nil
}
