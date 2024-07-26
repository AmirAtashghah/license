package hardware

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

type Fingerprint struct {
	MACAddress  []string `json:"mac_address"`
	UUID        string   `json:"uuid"`
	DiskSerial  string   `json:"disk_serial"`
	IPAddress   string   `json:"ip_address"`
	BIOSVersion string   `json:"bios_version"`
}

func GetMACAddress() ([]string, error) {
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

func GetUUID() (string, error) {
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

	SudoPassword := "amir5412"

	if _, err := stdin.Write([]byte(SudoPassword + "\n")); err != nil {
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

func GetDiskSerial() (string, error) {
	out, err := exec.Command("lsblk", "-o", "NAME,SERIAL").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func GetIPAddress() (string, error) {
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

func GetBIOSVersion() (string, error) {
	out, err := exec.Command("sudo", "dmidecode").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
