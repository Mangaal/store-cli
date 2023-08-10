/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all uploaded files in server",
	Long: `List all uploaded files in server
	        Example:
			
			store ls
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		ListFile()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListFile() (string, error) {

	// Make the HTTP POST request
	url := URL + "/apis/files"
	response, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return "", err

	}

	defer resp.Body.Close()

	res, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {

		fmt.Println("Error from server:", resp.Status)

		return "", err
	}

	GoData := struct {
		Fies []string `json:"files"`
	}{}

	json.Unmarshal(res, &GoData)

	for _, fname := range GoData.Fies {

		fmt.Println(fname)
	}

	return string(res), nil

}
