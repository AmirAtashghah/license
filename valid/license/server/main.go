package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

// give it from config file with unanimous name
var instanceFingerprint string
var sudoPassword string

type Fingerprint struct {
	MACAddress  []string `json:"mac_address"`
	UUID        string   `json:"uuid"`
	DiskSerial  string   `json:"disk_serial"`
	IPAddress   string   `json:"ip_address"`
	BIOSVersion string   `json:"bios_version"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	instanceFingerprint = os.Getenv("SYSTEM_INFO")

	// todo redis save random number
}

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

func createFingerprint() (bool, error) {

	mac, err := getMACAddress()
	if err != nil {
		return false, err
	}

	uuid, err := getUUID()
	if err != nil {
		return false, err
	}

	diskSerial, err := getDiskSerial()
	if err != nil {
		return false, err
	}

	ip, err := getIPAddress()
	if err != nil {
		return false, err
	}

	bios, err := getBIOSVersion()
	if err != nil {
		return false, err
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
		return false, err
	}

	encrypted := hashInfo(string(jsonInfo))

	if encrypted != instanceFingerprint {
		return false, nil
	}

	return true, nil
}

func hashInfo(input string) string {
	hash := sha256.New()

	hash.Write([]byte(input))

	hashedBytes := hash.Sum(nil)

	hashedString := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashedString
}

type checkBody struct {
	RandomNum string `json:"random_num"`
	Timestamp int64  `json:"timestamp"`
}

func check(c *fiber.Ctx) error {

	cb := new(checkBody)

	if err := json.Unmarshal(c.Body(), cb); err != nil {
		return err
	}

	// check time stamp duration
	timestamp := time.Unix(cb.Timestamp, 0)

	currentTime := time.Now()

	maxDuration := 5 * time.Minute

	if timestamp.After(currentTime) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Timestamp is in the future",
		})
	}

	if currentTime.Sub(timestamp) > maxDuration {
		return c.Status(400).JSON(fiber.Map{
			"error": "Timestamp is more than 5 minutes old",
		})
	}

	// todo check random number

	// todo store random number

	match, err := createFingerprint()
	if err != nil {
		log.Fatalf("Failed to check fingerprint: %v", err)
	}

	if !match {
		log.Fatal("Mismatch detected.")
	}

	return c.Status(200).JSON(fiber.Map{"message": "ok"})
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <sudo_password>")
		return
	}

	sudoPassword = os.Args[1]

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Post("/check-license", check)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
