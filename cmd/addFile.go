/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addFileCmd represents the addFile command
var addFileCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addFile called")

		sendFile(args)

	},
}

func init() {
	rootCmd.AddCommand(addFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func sendFile(files []string) {

	// Prepare the form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add multiple files to the form data // Add your file paths here
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		fileField, err := writer.CreateFormFile("files", filepath.Base(filePath))
		if err != nil {
			fmt.Println("Error creating form file:", err)
			return
		}
		_, err = io.Copy(fileField, file)
		if err != nil {
			fmt.Println("Error copying file data:", err)
			return
		}
	}

	// Close the form data
	writer.Close()

	// Make the HTTP POST request
	url := "http://" + URL + "/file"
	response, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer response.Body.Close()

	// Process the response
	fmt.Println("Response status:", response.Status)

}
