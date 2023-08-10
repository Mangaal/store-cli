package cmd_test

import (
	"net/http"
	"net/http/httptest"
	"store-cli/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Files uploaded successfully"}`))
	}))
	cmd.Data_Dir = "./test"
	cmd.URL = testServer.URL

	files := []string{}

	files = append(files, "./test/test.txt")

	errors, message := cmd.File(files)

	if errors != nil {
		t.Error()
	}

	assert.Contains(t, message, "Files uploaded successfully")

}

func TestSameFiles(t *testing.T) {

	cmd.Data_Dir = "./test"

	files := []string{}

	files = append(files, "./test/test.txt")

	errors, message := cmd.File(files)

	if errors != nil {
		t.Error()
	}

	assert.Contains(t, message, "Files in sync no change")

}

func TestCheck(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Files uploaded successfully"}`))
	}))
	cmd.Data_Dir = "./test"
	cmd.URL = testServer.URL

	var isUpdate *bool
	value := false
	isUpdate = &value

	isTrue := cmd.Check("./test/same.txt", "apple", isUpdate)

	assert.Equal(t, isTrue, false)
	assert.Equal(t, *isUpdate, true)

}
