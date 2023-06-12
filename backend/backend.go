package backend

import (
	"fmt"
	"os"
	"path/filepath"
)

func Initialize() {
	appDataDir, err := getAppDataDir()
	if err != nil {
		fmt.Printf("Failed to get app data directory: %s\n", err.Error())
		return
	}

	configDir := filepath.Join(appDataDir, "miranda-browser")
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return
	}

	configFile := filepath.Join(appDataDir, "mirandaconfig.conf")
	if !fileExists(configFile) {
		err := createFile(configFile)
		if err != nil {
			fmt.Printf("Failed to create config file: %s\n", err.Error())
		} else {
			fmt.Println("Config file created.")
		}
	}

	historyFile := filepath.Join(appDataDir, "mirandabrowserhistory.conf")
	if !fileExists(historyFile) {
		err := createFile(historyFile)
		if err != nil {
			fmt.Printf("Failed to create history file: %s\n", err.Error())
		} else {
			fmt.Println("History file created.")
		}
	}

	statusFile := filepath.Join(appDataDir, "mirandabrowsertabstatus.conf")
	if !fileExists(statusFile) {
		err := createFile(statusFile)
		if err != nil {
			fmt.Printf("Failed to create status file: %s\n", err.Error())
		} else {
			fmt.Println("Status file created.")
		}
	}
}

func getAppDataDir() (string, error) {
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(appDataDir, "miranda-browser"), nil
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func createFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
