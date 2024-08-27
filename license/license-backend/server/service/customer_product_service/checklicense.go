package customer_product_service

import (
	"encoding/json"
	"fmt"
	"server/pkg/encrypt"
	"server/pkg/hash"
	"server/pkg/param"
	"time"
)

// GenerateAuthKey :generate key with timpanist of request and random number do some think on these two values and send it to postgresqlclient
func (p Service) GenerateAuthKey(timestamp, number int64) (string, error) {

	authKey := fmt.Sprintf("%d%d", timestamp, number)

	authKeyHash, err := hash.Hash(authKey)
	if err != nil {

		return "", err
	}

	return authKeyHash, nil
}

func (p Service) ValidateClientHashInfo(id, hardwareHash string) (bool, error) {

	cp, err := p.repo.GetByID(id)
	if err != nil {

		return false, err
	}

	if cp == nil {
		return false, nil
	}

	if cp.HardwareHash != hardwareHash {
		return false, nil
	}

	return true, nil
}

func (p Service) ValidateTimestamp(timestamp int64) error {

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

func (p Service) CheckTimesConditions(ExpireAt, FirstConfirmedTime, LastConfirmedTime, reqTime int64) bool {
	currentTime := time.Now().Unix()

	if ExpireAt != 0 && ExpireAt < reqTime {
		//log.Println("ExpireAt != 0 && ExpireAt < reqTime ")
		return false
	}
	if FirstConfirmedTime != 0 && FirstConfirmedTime > reqTime {
		//log.Println("FirstConfirmedTime != 0 && FirstConfirmedTime > reqTime")

		return false
	}
	if LastConfirmedTime != 0 && LastConfirmedTime > reqTime {
		//log.Println("LastConfirmedTime != 0 && LastConfirmedTime > reqTime")

		return false
	}
	if currentTime < reqTime {
		//log.Printf("currentTime %d > reqTime %d ", currentTime, reqTime)

		return false
	}

	return true
}

//
//func (s Service) CheckDuplicateRequests(number int64) error {
//
//	cachedNumber, err := s.cache.GetRequestNumber(strconv.FormatInt(number, 10))
//	if err != nil {
//		logger.L().WithGroup(group).Error("error", "error", err.Error())
//
//		return err
//	}
//	logger.L().Info("msg", "cachedNumber", cachedNumber, "key", number)
//
//	if cachedNumber != "" {
//		logger.L().WithGroup(group).Error("mybug", "error", err.Error())
//
//		return fmt.Errorf("duplicate request")
//	}
//	logger.L().Info("here", "number", number)
//
//	err = s.cache.CacheRequestNumber(strconv.FormatInt(number, 10), strconv.FormatInt(number, 10))
//	if err != nil {
//		logger.L().WithGroup(group).Error("error", "error", err.Error())
//
//		return err
//	}
//
//	return nil
//}

func (p Service) EncryptResponse(response param.CheckLicenseResponse) (string, error) {

	byteResponse, err := json.Marshal(response)
	if err != nil {

		return "", err
	}

	encryptResponse := encrypt.Encrypt(string(byteResponse))

	return encryptResponse, nil
}
