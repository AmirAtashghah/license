package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

var (
	secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func Encrypt(text string) string {

	//aes, err := aes.NewCipher([]byte(secretKey))
	//if err != nil {
	//	panic(err)
	//}
	//
	//gcm, err := cipher.NewGCM(aes)
	//if err != nil {
	//	panic(err)
	//}
	//
	//nonce := make([]byte, gcm.NonceSize())
	//_, err = rand.Read(nonce)
	//if err != nil {
	//	panic(err)
	//}
	//
	//ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	fmt.Println("Encryption Program v0.01")

	//text := []byte("My Super Secret Code Stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	//fmt.Println(gcm.Seal(nonce, nonce, text, nil))

	return string(gcm.Seal(nonce, nonce, []byte(text), nil))
}

func Decrypt(ciphertext string) []byte {
	//fmt.Println("---------------------------------------------")
	//ciphertext = "13c55e6f245c348fb92779841fc52b8971ad2bda641d48e8b03f13290f56b24be7b4e5047779714336038373b5895ef140d360c302488ede3848c9722f548167ada2921e4b5082b4863f5dd7163d0fc50b669f27c997187646ebfb9d356b425c55148aa299f906664938e5f23990fdc8c5a11baba22cae5fa24e571d0faec8c38e651e4b2940d945350e36c41eee125480682d24367fea9fa83574e502580978adaec23a0a9a928bc6afe8524b9e9541ef59dcc8aeea3a566147a2d8ffbf53d20afa6d78f3238825be1da981548f94d1a49303"
	//
	//aes, err := aes.NewCipher([]byte(secretKey))
	//if err != nil {
	//	panic(err)
	//}
	//
	//gcm, err := cipher.NewGCM(aes)
	//if err != nil {
	//	panic(err)
	//}
	//
	//nonceSize := gcm.NonceSize()
	//nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	//
	//plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//return plaintext

	fmt.Println("Decryption Program v0.01")

	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		fmt.Println(err)
	}
	return plaintext
}
