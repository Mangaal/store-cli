package cmd

import (
	"os"
)

var URL string

func InitURL() {

	URL = os.Getenv("STORE_URL")
}

var Data_Dir string

func UploadDirectoryExists(uploadDirectory string) error {
	err := os.MkdirAll(uploadDirectory, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := os.Stat(uploadDirectory + "/data.json"); os.IsNotExist(err) {

		os.Create(uploadDirectory + "/data.json")

	}

	return nil
}
