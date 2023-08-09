package cmd

import "os"

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
	return nil
}
