package customer_product_service

import (
	"github.com/google/uuid"
	"server/entity"
	"server/pkg/param"
	"time"
)

type CustomerProductRepo interface {
	Insert(customerProduct *entity.CustomerProduct) error
	Update(customerProduct *entity.CustomerProduct) error
	Delete(id string) error
	GetByID(id string) (*entity.CustomerProduct, error)

	GetAll(limit, offset int16) ([]*entity.CustomerProduct, error)
	GetByCustomerIDAndProductID(customerID string, productID string) (*entity.CustomerProduct, error)
	GetByProductID(productID string) (*entity.CustomerProduct, error)
	GetByCustomerID(customerID string) (*entity.CustomerProduct, error)
}

type Service struct {
	repo CustomerProductRepo
}

func New(repo CustomerProductRepo) *Service {
	return &Service{repo: repo}
}

func (p Service) AddCustomerProduct(req *param.CreateCustomerProductRequest) (string, error) {

	customerProduct := &entity.CustomerProduct{
		ID:               uuid.New().String(),
		CustomerID:       req.CustomerID,
		ProductID:        req.ProductID,
		HardwareHash:     "hash",
		LicenseType:      req.LicenseType,
		IsActive:         *req.IsActive,
		ExpireAt:         req.ExpireAt,
		FirstConfirmedAt: -1,
		LastConfirmedAt:  -1,
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        -1,
	}

	if err := p.repo.Insert(customerProduct); err != nil {

		return "", err
	}

	return customerProduct.ID, nil
}

func (p Service) UpdateCustomerProduct(req *param.UpdateCustomerProductRequest) error {

	customerProduct := &entity.CustomerProduct{
		ID:               req.ID,
		CustomerID:       req.CustomerID,
		ProductID:        req.ProductID,
		HardwareHash:     req.HardwareHash,
		LicenseType:      req.LicenseType,
		IsActive:         *req.IsActive,
		ExpireAt:         req.ExpireAt,
		FirstConfirmedAt: req.FirstConfirmedAt,
		LastConfirmedAt:  req.LastConfirmedAt,
		CreatedAt:        req.CreatedAt,
		UpdatedAt:        time.Now().Unix(),
	}

	if err := p.repo.Update(customerProduct); err != nil {

		return err
	}

	return nil
}

func (p Service) DeleteCustomerProduct(req *param.DeleteCustomerProductRequest) error {

	if err := p.repo.Delete(req.ID); err != nil {

		return err
	}

	return nil
}

func (p Service) GetCustomerProductByID(req *param.GetCustomerProductRequest) (*entity.CustomerProduct, error) {

	customerProduct, err := p.repo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return customerProduct, nil
}

func (p Service) GetAll(filters *param.GetAllCustomerProductsRequest) ([]*entity.CustomerProduct, error) {

	customerProducts, err := p.repo.GetAll(filters.Limit, filters.Offset)
	if err != nil {

		return nil, err
	}

	return customerProducts, nil
}

func (p Service) CheckExist(productID, customerID string) (bool, error) {

	customerProduct, err := p.repo.GetByCustomerIDAndProductID(customerID, productID)
	if err != nil {
		return false, err
	}

	if customerProduct == nil {
		return false, nil
	}

	return true, nil
}

func (p Service) CheckExistByCustomerID(customerID string) (bool, error) {

	customerProduct, err := p.repo.GetByCustomerID(customerID)
	if err != nil {
		return false, err
	}

	if customerProduct == nil {
		return false, nil
	}

	return true, nil
}

func (p Service) CheckExistByProductID(productID string) (bool, error) {

	customerProduct, err := p.repo.GetByProductID(productID)
	if err != nil {
		return false, err
	}

	if customerProduct == nil {
		return false, nil
	}

	return true, nil
}
