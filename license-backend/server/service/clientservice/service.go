package clientservice

import (
	"fmt"
	"server/entity"
	"server/pkg/hash"
	"server/pkg/param"
	"strconv"
	"time"
)

type ClientRepo interface {
	Insert(client entity.Client) error
	Update(client entity.Client) error
	Delete(id string) error
	GetByID(id string) (entity.Client, error)
	GetAll(filter param.ClientFilter) ([]entity.Client, error)
}

type ClientCache interface {
	GetRequestNumber(number string) (string, error)
	CacheRequestNumber(number, key string) error
}

type Service struct {
	repo  ClientRepo
	cache ClientCache
}

func (s Service) CheckDuplicateRequests(number int64) error {

	cachedNumber, err := s.cache.GetRequestNumber(strconv.FormatInt(number, 10))
	if err != nil {
		return err
	}

	if cachedNumber != "" {
		return fmt.Errorf("duplicate request")
	}

	err = s.cache.CacheRequestNumber(strconv.FormatInt(number, 10), strconv.FormatInt(number, 10))
	if err != nil {
		return err
	}

	return nil
}

// GenerateAuthKey :generate key with timpanist of request and random number do some think on these two values and send it to postgresqlclient
func (s Service) GenerateAuthKey(timestamp, number int64) (string, error) {

	authKey := fmt.Sprintf("%d%d", timestamp, number)

	authKeyHash, err := hash.Hash(authKey)
	if err != nil {
		return "", err
	}
	return authKeyHash, nil
}

func (s Service) AddNewClient(req *param.ClientRequest) error {

	timestamp := time.Now().Unix()

	// todo change value if need ?
	c := entity.Client{
		ID:              req.ID,
		SoftwareName:    &req.SoftwareName,
		SoftwareVersion: &req.SoftwareVersion,
		HardwareHash:    &req.HardwareHash,
		LicenseType:     nil,
		UserMetadata:    &req.UserMetadata,
		ExpiresAt:       nil,
		CreatedAt:       &timestamp,
		UpdatedAt:       nil,
		DeletedAt:       nil,
	}

	if err := s.repo.Insert(c); err != nil {
		return err
	}

	return nil
}

func (s Service) ValidateClientHashInfo(id, hardwareHash string) (bool, error) {

	client, err := s.repo.GetByID(id)
	if err != nil {
		return false, err
	}

	if *client.HardwareHash != hardwareHash {
		return false, nil
	}

	return true, nil
}

func (s Service) ValidateTimestamp(timestamp int64) error {

	tsp := time.Unix(timestamp, 0)

	currentTime := time.Now()

	// Define the maximum allowable duration (5 minutes) todo change duration if need
	maxDuration := 5 * time.Minute

	if tsp.After(currentTime) {
		return fmt.Errorf("invalid request")
	}

	if currentTime.Sub(tsp) > maxDuration {
		return fmt.Errorf("invalid request")
	}

	return nil
}

// todo implement

func (s Service) DeleteClient(id string) error {

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s Service) UpdateClient(req param.UpdateClientRequest) error {

	timestamp := time.Now().Unix()

	c := entity.Client{
		ID:           req.ID,
		LicenseType:  &req.LicenseType,
		UserMetadata: &req.UserMetadata,
		ExpiresAt:    &req.ExpiresAt,
		UpdatedAt:    &timestamp,
		DeletedAt:    nil,
	}

	if err := s.repo.Update(c); err != nil {
		return err
	}

	return nil
}

func (s Service) ListClients(filter param.ClientFilter) ([]entity.Client, error) {

	if *filter.Offset < 0 {
		*filter.Offset = 0
	}
	if *filter.Limit <= 0 {
		*filter.Limit = 20
	}

	clients, err := s.repo.GetAll(filter)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (s Service) ChangeActiveStatus(req param.ChangeActivateRequest) error {

	timestamp := time.Now().Unix()

	c := entity.Client{
		ID:        req.ID,
		IsActive:  &req.IsActivate,
		UpdatedAt: &timestamp,
	}

	if err := s.repo.Update(c); err != nil {
		return err
	}

	return nil
}
