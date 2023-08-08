/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// optionCmd represents the option command
var optionCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("option called")

		limit, _ := cmd.Flags().GetString("limit")

		sort, _ := cmd.Flags().GetString("sort")

		options(limit, sort)

	},
}

func init() {
	rootCmd.AddCommand(optionCmd)

	// Here you will define your flags and configuration settings.

	//optionCmd.PersistentFlags().String("limit", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	optionCmd.Flags().StringP("limit", "l", "10", "Help message for toggle")
	optionCmd.Flags().StringP("sort", "s", "d", "Help message for toggle")
}

func options(limit string, sort string) {

	// Make the HTTP POST request
	url := "http://" + URL + "/file/option/" + sort + "/" + limit
	response, err := http.NewRequest("GET", url, nil)
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

}
