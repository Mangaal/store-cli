/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete Files from Server ",
	Long: `Delete Files from Server 
	         Example:

	         store rm filenames....
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		if len(args) == 0 {

			fmt.Println("got empty argument")

			return
		}
		deleteFile(args)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type FileBody struct {
	Files []string `json:"files"`
}

func deleteFile(files []string) {

	FileBodyreq := FileBody{}

	FileBodyreq.Files = append(FileBodyreq.Files, files...)

	body, _ := json.Marshal(FileBodyreq)

	// Make the HTTP POST request
	url := "http://" + URL + "/file/"
	response, err := http.NewRequest("DELETE", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	res, _ := io.ReadAll(resp.Body)

	// Process the response
	fmt.Println("Response status:", string(res))

}
