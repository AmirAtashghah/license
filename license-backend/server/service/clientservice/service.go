package clientservice

import (
	"fmt"
	"server/entity"
	"server/pkg/param"
	"strconv"
	"time"
)

type ClientRepo interface {
	Insert(client entity.Client) error
	Update(client entity.Client) error
	Delete(client entity.Client) error
	GetById(id string) (entity.Client, error)
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

	authKeyHash, err := hashPassword(authKey)
	if err != nil {
		return "", err
	}
	return authKeyHash, nil
}

func (s Service) AddNewClient(req *param.ClientRequest) error {

	// todo change value if need ?
	c := entity.Client{
		ID:              req.ID,
		SoftwareName:    req.SoftwareName,
		SoftwareVersion: req.SoftwareVersion,
		HardwareHash:    req.HardwareHash,
		LicenseType:     "",
		UserMetadata:    req.UserMetadata,
		ExpiresAt:       0,
		CreatedAt:       time.Now().Unix(),
		UpdatedAt:       0,
		DeletedAt:       0,
	}

	if err := s.repo.Insert(c); err != nil {
		return err
	}

	return nil
}

func (s Service) ValidateClientHashInfo(id, hardwareHash string) (bool, error) {

	client, err := s.repo.GetById(id)
	if err != nil {
		return false, err
	}

	if client.HardwareHash != hardwareHash {
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

func (s Service) DeleteClient() {}

func (s Service) UpdateClient() {}

func (s Service) ListClients() {}
