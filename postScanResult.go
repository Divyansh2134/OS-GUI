package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/models"
)

func postScanResult(scanResult models.ScanResult, config models.Config) error {

	jsonData, err := json.Marshal(scanResult)
	if err != nil {
		log.Println("Failed to marshal scan result:", err)
		return err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", config.ServerURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Failed to create request:", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to post scan result:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to post scan result, status code: %d\n", resp.StatusCode)
		return fmt.Errorf("failed to post scan result, status code: %d", resp.StatusCode)
	}

	log.Println("Scan result posted successfully!")
	return nil
}
