/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"store-cli/cmd"
)

func main() {

	cmd.Data_Dir = os.Getenv("DATA_DIR")
	cmd.UploadDirectoryExists(cmd.Data_Dir)
	cmd.InitURL()
	cmd.Execute()
}
