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

// optionCmd represents the option command
var optionCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "Get most frequently use words from the uploaded documentes",
	Long: `Get most frequently use words from the uploaded documentes
	         Example:

	         store freq-words
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("option called")

		limit, _ := cmd.Flags().GetString("limit")

		sort, _ := cmd.Flags().GetString("sort")

		Options(limit, sort)

	},
}

func init() {
	rootCmd.AddCommand(optionCmd)

	// Here you will define your flags and configuration settings.

	//optionCmd.PersistentFlags().String("limit", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	optionCmd.Flags().StringP("limit", "l", "10", "limit the no of output")
	optionCmd.Flags().StringP("sort", "s", "d", "to get output stort in ascending or descending pass a for assending and b for decending")
}

func Options(limit string, sort string) (string, error) {

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

	if resp.StatusCode != 200 {

		fmt.Println("Error making POST request:", err)
		return "", err
	}

	defer resp.Body.Close()

	res, _ := io.ReadAll(resp.Body)

	var WordFrequencyList struct {
		Items []struct {
			Word      string `json:"word"`
			Frequency int    `json:"frequency"`
		} `json:"items"`
	}

	json.Unmarshal(res, &WordFrequencyList)

	for _, fname := range WordFrequencyList.Items {

		fmt.Printf("%d   %s \n", fname.Frequency, fname.Word)
	}

	return string(res), nil

}
