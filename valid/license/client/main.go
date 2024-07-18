package main

import (
	"crypto/rand"
	"math/big"
)

func main() {

	// give hardware info 1. UUID, 2. Mac address, 3. Disk serial, 4. Cpu info, 5. IP address, 6. BIOS version, 7. Host or Domain name

	// generate random number

	// generate time stamp

	// encrypt body of request by private key

	// call validation api set body and uuid as param

	// give response decrypt

	// check data that generate by random number and time stamp

}

func generateRandomNumber() (string, error) {

	min := 1000
	max := 10000
	rangeSize := max - min + 1
	nInRange, err := rand.Int(rand.Reader, big.NewInt(int64(rangeSize)))
	if err != nil {
		return "", err
	}

	nInRange.Add(nInRange, big.NewInt(int64(min)))

	return nInRange.String(), nil
}
