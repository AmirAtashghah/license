package apicaller

import (
	"bytes"
	"client/hardware"
	"client/utility"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type IdRepo interface {
	Get(id int) (string, error)
	Set(clientID string) error // todo set client id by id 1. there is one record in database table id:1,clientID:UUID
}

type Service struct {
	repo IdRepo
}

func NewService(repo IdRepo) *Service {
	return &Service{repo: repo}
}

type bodyRequest struct {
	ID              string `json:"id"`
	SoftwareName    string `json:"software_name" validate:"required"`
	SoftwareVersion string `json:"software_version" validate:"required"`
	HardwareHash    string `json:"hardware_hash" validate:"required,sha256"`
	LicenseType     string `json:"license_type" validate:"required"`
	UserMetadata    string `json:"user_metadata" validate:"required"`
	TimeStamp       int64  `json:"time_stamp" validate:"required,numeric"`
	RandomNumber    int64  `json:"random_number" validate:"required,numeric"`
}

type bodyResponse struct {
	ClientID         string `json:"client_id"`
	AuthKey          string `json:"auth_key"`
	ValidationStatus bool   `json:"validation_status"`
}

func createFingerprint() (string, error) {

	mac, err := hardware.GetMACAddress()
	if err != nil {
		return "", err
	}

	uuid, err := hardware.GetUUID()
	if err != nil {
		return "", err
	}

	diskSerial, err := hardware.GetDiskSerial()
	if err != nil {
		return "", err
	}

	ip, err := hardware.GetIPAddress()
	if err != nil {
		return "", err
	}

	bios, err := hardware.GetBIOSVersion()
	if err != nil {
		return "", err
	}

	currentFingerprint := hardware.Fingerprint{
		MACAddress:  mac,
		UUID:        uuid,
		DiskSerial:  diskSerial,
		IPAddress:   ip,
		BIOSVersion: bios,
	}

	jsonInfo, err := json.Marshal(currentFingerprint)
	if err != nil {
		return "", err
	}

	hash := sha256.New()

	hash.Write(jsonInfo)

	hashedBytes := hash.Sum(nil)

	encryptedSystemInfo := base64.StdEncoding.EncodeToString(hashedBytes)

	return encryptedSystemInfo, nil
}

func (s Service) Check() error {
	timestamp := time.Now().Unix()

	randomNumber, err := utility.GenerateRandomNumber()
	if err != nil {
		return err
	}

	hash, err := createFingerprint()
	if err != nil {
		return err
	}

	// todo set defaults value : "" . this case happen when client run first time and it does not have id
	clientID, err := s.repo.Get(1)
	if err != nil {
		return err
	}

	// todo complete values
	b := bodyRequest{
		ID:              clientID,
		SoftwareName:    "citra",
		SoftwareVersion: "V1.0.0",
		HardwareHash:    hash,
		LicenseType:     "test",
		UserMetadata:    "test",
		TimeStamp:       timestamp,
		RandomNumber:    randomNumber,
	}

	jsonData, err := json.Marshal(b)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:3000/check-license", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyRes := new(bodyResponse)

	if resp.StatusCode == 200 {

		if err := json.Unmarshal(body, bodyRes); err != nil {
			return err
		}

		clientAuthKey := fmt.Sprintf("%d%d", timestamp, randomNumber)

		if bodyRes.AuthKey != "" {
			if !utility.CheckAuthHash(clientAuthKey, bodyRes.AuthKey) {
				log.Println("invalid server auth key")
				os.Exit(1)
				// todo stop app
			}
		}

		if !bodyRes.ValidationStatus {
			log.Println("invalid client")
			os.Exit(1)
			// todo stop app
		}

		if bodyRes.ClientID != "" {
			if err := s.repo.Set(bodyRes.ClientID); err != nil {
				return err
			}
			// todo in this case call check license again with new id ?
		}

		return nil

	} else {
		return fmt.Errorf(string(body))
		// todo do something when for example 5 time request failed
	}
}
