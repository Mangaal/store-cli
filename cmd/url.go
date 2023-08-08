package cmd

import "os"

var URL string

func InitURL() {

	URL = os.Getenv("STORE_URL")
}
