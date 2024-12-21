package main

import (
	"fmt"
	"log"

	"example.com/m/models"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const configFile = "config.json"

func main() {
	application := app.New()
	window := application.NewWindow("Golang Desktop App")

	var os models.ScanResult
	config := loadConfig()

	serverURLLabel := widget.NewLabel("Server URL:")
	serverURL := widget.NewEntry()
	serverURL.SetText(config.ServerURL)
	apiKeyLabel := widget.NewLabel("API Key:")
	apiKey := widget.NewEntry()
	apiKey.SetText(config.APIKey)

	saveButton := widget.NewButton("Save", func() {
		go func() {
			config.ServerURL = serverURL.Text
			config.APIKey = apiKey.Text
			saveConfig(config)
			fmt.Println("Configuration saved!")
		}()
	})

	scanButton := widget.NewButton("Scan Now", func() {
		go func() {
			os = performScan()
			err := postScanResult(os, config)
			if err != nil {
				log.Println("Error posting scan result:", err)
			}
			fmt.Println("Scan result:", os)
		}()
	})

	content := container.NewVBox(
		widget.NewLabel("Welcome to Golang Desktop App"),
		serverURLLabel,
		serverURL,
		apiKeyLabel,
		apiKey,
		saveButton,
		scanButton,
	)
	window.SetContent(content)
	window.ShowAndRun()
}
