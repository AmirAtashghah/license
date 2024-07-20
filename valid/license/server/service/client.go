package service

import (
	"server/pkg/param"
)

func CheckDuplicateRequests(number string) error {
	return nil
}

// GenerateAuthKey :generate key with timpanist of request and random number do some think on these two values and send it to client
func GenerateAuthKey(timestamp, number int64) (string, error) {

	return "", nil
}

func AddNewClient(client *param.ClientRequest) error {
	return nil
}

func DeleteClient() {}

func ValidateClientHashInfo(hardwareHash string) error {
	return nil
}

func ValidateTimestamp(timestamp int64) error {
	return nil
}

func ValidateDuplicateRequest() {}
