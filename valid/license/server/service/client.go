package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/db"
	"server/entity"
	"server/param"
	"server/repository"
	"strconv"
	"time"
)

func CheckDuplicateRequests(number int64) error {

	cachedNumber, err := repository.GetCachedNumber(strconv.FormatInt(number, 10))
	if err != nil {
		return err
	}

	if cachedNumber != "" {
		return fmt.Errorf("duplicate request")
	}

	err = repository.CacheNumber(strconv.FormatInt(number, 10), strconv.FormatInt(number, 10))
	if err != nil {
		return err
	}

	return nil
}

// GenerateAuthKey :generate key with timpanist of request and random number do some think on these two values and send it to client
func GenerateAuthKey(timestamp, number int64) (string, error) {

	authKey := fmt.Sprintf("%d%d", timestamp, number)

	authKeyHash, err := hashPassword(authKey)
	if err != nil {
		return "", err
	}
	return authKeyHash, nil
}

func AddNewClient(req *param.ClientRequest) error {

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

	if err := repository.InsertClient(db.Conn, c); err != nil {
		return err
	}

	return nil
}

func ValidateClientHashInfo(id, hardwareHash string) (bool, error) {

	client, err := repository.GetClientByID(db.Conn, id)
	if err != nil {
		return false, err
	}

	if client.HardwareHash != hardwareHash {
		return false, nil
	}

	return true, nil
}

func ValidateTimestamp(timestamp int64) error {

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

func DeleteClient() {}

func UpdateClient() {}

func ListClients() {}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
