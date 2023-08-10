/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addFileCmd represents the addFile command
var addFileCmd = &cobra.Command{
	Use:   "files",
	Short: "Uploads files to server",
	Long: `Uploads files to server.
	         Example:

		     store files filename....
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addFile called")

		if len(args) == 0 {

			fmt.Println("got empty argument")

			return
		}

		File(args)

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

func File(files []string) (error, string) {

	// Prepare the form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	goData := map[string]string{}

	requestBodylen := requestBody.Len()

	var isUpdate *bool
	value := false
	isUpdate = &value

	// Add multiple files to the form data // Add your file paths here
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return err, ""
		}

		filetmp, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return err, ""
		}

		content, err := io.ReadAll(filetmp)

		if err != nil {

			return err, ""
		}

		if Check(file.Name(), string(content), isUpdate) {

			fileField, err := writer.CreateFormFile("files", filepath.Base(filePath))
			if err != nil {
				fmt.Println("Error creating form file:", err)
				return err, ""
			}
			_, err = io.Copy(fileField, file)
			if err != nil {
				fmt.Println("Error copying file data:", err)
				return err, ""
			}

			goData[file.Name()] = generateHash(string(content))

		}

	}

	if requestBody.Len() == requestBodylen && !*isUpdate {

		fmt.Println("Files in sync no change")

		return nil, "Files in sync no change"

	}

	if requestBody.Len() == requestBodylen && *isUpdate {

		fmt.Println("Files uploaded successfully")

		return nil, "Files uploaded successfully"

	}

	// Close the form data
	writer.Close()

	// Make the HTTP POST request
	url := URL + "/apis/file"
	response, err := http.Post(url, writer.FormDataContentType(), &requestBody)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return err, ""
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {

		fmt.Println("Error from server:", response.Status)

		return nil, ""
	}

	for name, key := range goData {

		updateDatabase(name, "", key)

	}

	body, _ := io.ReadAll(response.Body)
	// Process the response
	fmt.Println("Response status:", response.Status)

	return nil, string(body)

}

func Check(fname string, contant string, isupdate *bool) bool {

	file, err := os.Open(Data_Dir + "/data.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return true
	}
	defer file.Close()

	datajson, _ := io.ReadAll(file)
	goData := map[string]string{}

	json.Unmarshal(datajson, &goData)

	for dkey, code := range goData {

		if fname != dkey && code == generateHash(contant) {

			// Make the HTTP POST request
			url := URL + "/apis/file/" + fname + "/" + dkey
			response, err := http.NewRequest("POST", url, nil)
			if err != nil {
				fmt.Println("Error making POST request:", err)
				return true
			}
			client := &http.Client{}
			resp, err := client.Do(response)
			if err != nil {
				log.Fatalln(err)
				return true
			}

			if resp.StatusCode != 200 {

				fmt.Println("Error from server:", resp.Status)

				return true
			}

			defer resp.Body.Close()

			res, _ := io.ReadAll(resp.Body)

			fmt.Println(string(res))

			updateDatabase(fname, dkey, code)
			*isupdate = true
			return false

		}
		if fname == dkey && code == generateHash(contant) {

			return false
		}

	}

	return true

}

func updateDatabase(new string, old string, value string) {

	file, err := os.Open(Data_Dir + "/data.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	datajson, _ := io.ReadAll(file)
	goData := map[string]string{}

	json.Unmarshal(datajson, &goData)

	goData[new] = value

	if old != "" {
		goData[old] = ""

	}

	datajson, _ = json.Marshal(goData)

	os.WriteFile(Data_Dir+"/data.json", datajson, 0644)

}

func generateHash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashValue := hasher.Sum(nil)
	return hex.EncodeToString(hashValue)
}
