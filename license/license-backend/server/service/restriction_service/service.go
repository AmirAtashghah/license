package restriction_service

import (
	"encoding/json"
	"log"
	"server/entity"
	"server/pkg/encrypt"
	"server/pkg/param"
	"strconv"
	"time"
)

type RestrictionRepo interface {
	Insert(restriction *entity.Restriction) error
	Update(restriction *entity.Restriction) error
	Delete(id int16) error
	GetByID(id string) (*entity.Restriction, error)
	//GetByProductID(productID string) ([]*entity.Restriction, error)
	GetByKey(key string) (*entity.Restriction, error)
	GetAll(filter *param.GetAllRestrictionsRequest) ([]*entity.Restriction, error)

	InsertCPR(cpr *entity.CustomersProductRestriction) error
	UpdateCPR(cpr *entity.CustomersProductRestriction) error
	DeleteCPR(id string) error
	GetByIDCPR(id string) (*entity.CustomersProductRestriction, error)
	GetByCustomersProductIDCPR(customersProductID string) ([]*entity.CustomersProductRestriction, error)
	GetByCustomersProductIDAndRestrictionIDCPR(customersProductID, restrictionID string) (*entity.CustomersProductRestriction, error)
	GetAllCPR(filter *param.GetAllRestrictionsRequest) ([]*entity.CustomersProductRestriction, error)
}

type Service struct {
	repo RestrictionRepo
}

func New(repo RestrictionRepo) *Service {
	return &Service{repo: repo}
}

func (p Service) AddRestriction(req *param.CreateRestrictionRequest) error {

	restriction := &entity.Restriction{
		Key:       req.Key,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: -1,
	}

	if err := p.repo.Insert(restriction); err != nil {

		return err
	}

	return nil
}

func (p Service) UpdateCustomer(req *param.UpdateRestrictionRequest) error {

	restriction := &entity.Restriction{
		ID:        req.ID,
		Key:       req.Key,
		CreatedAt: req.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}

	if err := p.repo.Update(restriction); err != nil {

		return err
	}

	return nil
}

func (p Service) DeleteRestriction(req *param.DeleteRestrictionRequest) error {

	num, err := strconv.ParseInt(req.ID, 10, 16)
	if err != nil {
		return err
	}

	if err := p.repo.Delete(int16(num)); err != nil {

		return err
	}

	return nil
}

func (p Service) GetRestrictionByID(req *param.GetRestrictionRequest) (*entity.Restriction, error) {

	restriction, err := p.repo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return restriction, nil
}

//func (p Service) GetRestrictionsByProductID(req *param.GetRestrictionsByProductRequest) ([]*entity.Restriction, error) {
//
//	restrictions, err := p.repo.GetByProductID(req.ProductID)
//	if err != nil {
//
//		return nil, err
//	}
//
//	return restrictions, nil
//}

func (p Service) GetAllRestrictions(filter *param.GetAllRestrictionsRequest) ([]*entity.Restriction, error) {

	restrictions, err := p.repo.GetAll(filter)
	if err != nil {
		return nil, err
	}

	return restrictions, nil
}

func (p Service) CheckExist(key string) (bool, error) {

	restriction, err := p.repo.GetByKey(key)
	if err != nil {
		return false, err
	}

	if restriction == nil {
		return false, nil
	}

	return true, nil
}

func (p Service) EncryptRestriction(result any) (string, error) {

	byteRestriction, err := json.Marshal(result)
	if err != nil {

		return "", err
	}

	encryptRestriction := encrypt.Encrypt(string(byteRestriction))

	return encryptRestriction, nil

}

func (p Service) AddCustomersProductRestriction(req *param.CreateCustomersProductRestrictionRequest) error {

	log.Println("162 : ", req)

	var result []struct {
		RestrictionID int16  `json:"id"`
		Value         string `json:"value"`
	}

	if err := json.Unmarshal([]byte(req.RestrictionIDAndValues), &result); err != nil {
		return err
	}

	log.Println("result :", result)

	for _, restriction := range result {

		cpr := &entity.CustomersProductRestriction{
			RestrictionID:      restriction.RestrictionID,
			CustomersProductID: req.CustomersProductID,
			Value:              restriction.Value,
			CreatedAt:          time.Now().Unix(),
			UpdatedAt:          -1,
		}

		if err := p.repo.InsertCPR(cpr); err != nil {

			return err
		}
	}

	return nil
}

func (p Service) DeleteCustomersProductRestriction(req *param.DeleteCustomersProductRestrictionRequest) error {

	if err := p.repo.DeleteCPR(req.CustomersProductID); err != nil {

		return err
	}

	return nil
}

func (p Service) GetCustomersProductRestrictionByID(req *param.GetCustomersProductRestrictionRequest) (*entity.CustomersProductRestriction, error) {

	cpr, err := p.repo.GetByIDCPR(req.ID)
	if err != nil {
		return nil, err
	}

	return cpr, nil
}

func (p Service) GetCustomersProductRestrictionByCustomersProductIDAndRestrictionID(req *param.GetCustomersProductRestrictionsByCPAndRestrictionIDRequest) (*entity.CustomersProductRestriction, error) {

	cprs, err := p.repo.GetByCustomersProductIDAndRestrictionIDCPR(req.CustomersProductID, req.RestrictionID)
	if err != nil {

		return nil, err
	}

	return cprs, nil
}

func (p Service) GetCustomersProductRestrictionByCustomersProductID(req *param.GetCustomersProductRestrictionsByCustomerProductIDRequest) ([]*entity.CustomersProductRestriction, error) {

	cprs, err := p.repo.GetByCustomersProductIDCPR(req.CustomersProductID)
	if err != nil {

		return nil, err
	}

	return cprs, nil
}

func (p Service) GetAllCustomersProductRestrictions(filter *param.GetAllRestrictionsRequest) ([]*entity.CustomersProductRestriction, error) {

	cprs, err := p.repo.GetAllCPR(filter)
	if err != nil {
		return nil, err
	}

	return cprs, nil
}

func (p Service) CheckExistCustomersProductRestriction(customersProductID, restrictionID string) (bool, error) {

	cpr, err := p.repo.GetByCustomersProductIDAndRestrictionIDCPR(customersProductID, restrictionID)
	if err != nil {
		return false, err
	}

	if cpr == nil {
		return false, nil
	}

	return true, nil
}
