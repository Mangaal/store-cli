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

// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "Counts the total no of words",
	Long: `Counts the total no of words
	         Example:

			store wc 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wc called")

		WC("1", "d")

	},
}

func init() {
	rootCmd.AddCommand(wcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func WC(limit string, sort string) (string, error) {

	// Make the HTTP POST request
	url := URL + "/apis/file/option/" + sort + "/" + limit
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

	var Count struct {
		TotalWordCount int `json:"totalWordCount"`
	}

	json.Unmarshal(res, &Count)

	fmt.Println("total_word_could ", Count.TotalWordCount)

	return string(res), nil

}
