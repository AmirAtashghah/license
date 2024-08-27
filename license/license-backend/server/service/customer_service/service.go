package customer_service

import (
	"github.com/google/uuid"
	"server/entity"
	"server/pkg/param"
	"time"
)

type CustomerRepo interface {
	Insert(customer *entity.Customer) error
	Update(customer *entity.Customer) error
	Delete(id string) error
	GetByID(id string) (*entity.Customer, error)
	GetAll(limit, offset int16) ([]*entity.Customer, error)
	GetCustomerByNameAndEmail(name, email string) (*entity.Customer, error)
}

type Service struct {
	repo CustomerRepo
}

func New(repo CustomerRepo) *Service {
	return &Service{repo: repo}
}

func (p Service) AddCustomer(req *param.CreateCustomerRequest) error {

	customer := &entity.Customer{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: -1,
	}

	if err := p.repo.Insert(customer); err != nil {

		return err
	}

	return nil
}

func (p Service) UpdateCustomer(req *param.UpdateCustomerRequest) error {

	customer := &entity.Customer{
		ID:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: req.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}

	if err := p.repo.Update(customer); err != nil {

		return err
	}

	return nil
}

func (p Service) DeleteCustomer(req *param.DeleteCustomerRequest) error {

	if err := p.repo.Delete(req.ID); err != nil {

		return err
	}

	return nil
}

func (p Service) GetCustomerByID(req *param.GetCustomerRequest) (*entity.Customer, error) {

	customer, err := p.repo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (p Service) GetAll(filters *param.GetAllCustomersRequest) ([]*entity.Customer, error) {

	customers, err := p.repo.GetAll(filters.Limit, filters.Offset)
	if err != nil {

		return nil, err
	}

	return customers, nil
}

func (p Service) CheckExist(name, email string) (bool, error) {

	customer, err := p.repo.GetCustomerByNameAndEmail(name, email)
	if err != nil {
		return false, err
	}

	if customer == nil {
		return false, nil
	}

	return true, nil
}
