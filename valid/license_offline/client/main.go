package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RequestData struct {
	Number    string `json:"number"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	// Create the data to send
	data := RequestData{
		Number:    "12345654",
		Timestamp: time.Now().Unix(),
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Send the POST request
	resp, err := http.Post("http://localhost:3000/check-license", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("Response:", result)
}
