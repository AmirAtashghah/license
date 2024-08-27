package hash

import (
	"crypto/sha512"
	"encoding/base64"
	"log"
)

//
//func Hash(str string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
//	return string(bytes), err
//}
//
//func CheckHash(str, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
//	return err == nil
//}

func Hash(password string) (string, error) {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	//passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash, nil
}

func CheckHash(currPassword, hashedPassword string) bool {

	log.Println("old hash", hashedPassword)

	var currPasswordHash, _ = Hash(currPassword)

	log.Println("new hash", currPasswordHash)
	return hashedPassword == currPasswordHash
}
