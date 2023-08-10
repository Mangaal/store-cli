package cmd_test

import (
	"net/http"
	"net/http/httptest"
	"store-cli/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteFile(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Files Deleted successfully"}`))
	}))
	cmd.Data_Dir = "./test"
	cmd.URL = testServer.URL

	files := []string{}

	files = append(files, "./test/test.txt")

	message, errors := cmd.DeleteFile(files)

	if errors != nil {
		t.Error()
	}

	assert.Contains(t, message, "Files Deleted successfully")

}
