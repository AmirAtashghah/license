package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

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

type Fingerprint struct {
	MACAddress  []string `json:"mac_address"`
	UUID        string   `json:"uuid"`
	DiskSerial  string   `json:"disk_serial"`
	IPAddress   string   `json:"ip_address"`
	BIOSVersion string   `json:"bios_version"`
}

var sudoPassword string

// todo save client id in database
var clientID string = "d1c6da2a-8ece-44d2-a5b1-20427b44287f"

// todo write scheduler for check license every * time
//func init() {
//
//}

//////// -----------

func getMACAddress() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}

	return as, nil
}

func getUUID() (string, error) {
	cmd := exec.Command("sudo", "-S", "cat", "/sys/class/dmi/id/product_uuid")

	var out bytes.Buffer
	cmd.Stdout = &out

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println("Error creating stdin pipe:", err)
		return "", err
	}

	if err := cmd.Start(); err != nil {
		log.Println("Error starting command:", err)
		return "", err
	}

	if _, err := stdin.Write([]byte(sudoPassword + "\n")); err != nil {
		log.Println("Error writing to stdin pipe:", err)
		return "", err
	}

	// Close the stdin pipe to signal the end of input
	if err := stdin.Close(); err != nil {
		log.Println("Error closing stdin pipe:", err)
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		log.Println("Error waiting for command:", err)
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}

func getDiskSerial() (string, error) {
	out, err := exec.Command("lsblk", "-o", "NAME,SERIAL").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func getIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no IP address found")
}

func getBIOSVersion() (string, error) {
	out, err := exec.Command("sudo", "dmidecode").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

/////// ----------

func createFingerprint() (string, error) {

	mac, err := getMACAddress()
	if err != nil {
		return "", err
	}

	uuid, err := getUUID()
	if err != nil {
		return "", err
	}

	diskSerial, err := getDiskSerial()
	if err != nil {
		return "", err
	}

	ip, err := getIPAddress()
	if err != nil {
		return "", err
	}

	bios, err := getBIOSVersion()
	if err != nil {
		return "", err
	}

	currentFingerprint := Fingerprint{
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

func check() error {

	timestamp := time.Now().Unix()

	randomNumber, err := generateRandomNumber()
	if err != nil {
		return err
	}

	hash, err := createFingerprint()
	if err != nil {
		return err
	}

	// todo complete values
	b := bodyRequest{
		ID:              clientID,
		SoftwareName:    "citra",
		SoftwareVersion: "V1.0.0",
		HardwareHash:    hash,
		LicenseType:     "aa",
		UserMetadata:    "aaa",
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
			if !checkAuthHash(clientAuthKey, bodyRes.AuthKey) {

				return fmt.Errorf("invalid server auth key")
				// todo stop app
			}
		}

		if !bodyRes.ValidationStatus {
			return fmt.Errorf("invalid client")
			// todo stop app
		}

		if bodyRes.ClientID != "" {
			clientID = bodyRes.ClientID
			// todo in this case call check license again with new id
		}

		return nil

	} else {
		return fmt.Errorf(string(body))
		// todo do something when for example 5 time request failed
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <sudo_password>")
		return
	}

	sudoPassword = os.Args[1]

	if err := check(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("checked license, every things ok")
}

func generateRandomNumber() (int64, error) {

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

func checkAuthHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
