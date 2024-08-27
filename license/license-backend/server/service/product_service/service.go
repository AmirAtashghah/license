package product_service

import (
	"github.com/google/uuid"
	"server/entity"
	"server/pkg/param"
	"time"
)

type ProductRepo interface {
	Insert(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id string) error
	GetByID(id string) (*entity.Product, error)
	GetAll(limit, offset int16) ([]*entity.Product, error)
	GetProductByNameAndVersion(name string, version string) (*entity.Product, error)
}

type Service struct {
	repo ProductRepo
}

func New(repo ProductRepo) *Service {
	return &Service{repo: repo}
}

func (p Service) AddNewProduct(req *param.CreateProductRequest) error {

	product := &entity.Product{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Title:     req.Title,
		Version:   req.Version,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: -1,
	}

	if err := p.repo.Insert(product); err != nil {

		return err
	}

	return nil
}

func (p Service) UpdateProduct(req *param.UpdateProductRequest) error {

	product := &entity.Product{
		ID:        req.ID,
		Name:      req.Name,
		Title:     req.Title,
		Version:   req.Version,
		CreatedAt: req.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}

	if err := p.repo.Update(product); err != nil {

		return err
	}

	return nil
}

func (p Service) DeleteProduct(req *param.DeleteProductRequest) error {

	if err := p.repo.Delete(req.ID); err != nil {

		return err
	}

	return nil
}

func (p Service) GetProductByID(req *param.GetProductRequest) (*entity.Product, error) {

	product, err := p.repo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p Service) GetAll(filters *param.GetAllProductsRequest) ([]*entity.Product, error) {

	products, err := p.repo.GetAll(filters.Limit, filters.Offset)
	if err != nil {

		return nil, err
	}

	return products, nil
}

func (p Service) CheckExist(name, version string) (bool, error) {

	product, err := p.repo.GetProductByNameAndVersion(name, version)
	if err != nil {
		return false, err
	}

	if product == nil {
		return false, nil
	}

	return true, nil
}
