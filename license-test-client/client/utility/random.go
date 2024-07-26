package utility

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomNumber() (int64, error) {

	mi := 1000
	ma := 10000
	rangeSize := ma - mi + 1
	nInRange, err := rand.Int(rand.Reader, big.NewInt(int64(rangeSize)))
	if err != nil {
		return 0, err
	}

	nInRange.Add(nInRange, big.NewInt(int64(mi)))

	return nInRange.Int64(), nil
}
