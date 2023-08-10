package cmd_test

import (
	"net/http"
	"net/http/httptest"
	"store-cli/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFile(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"files": ["new.yaml","new1.txt"] }`))
	}))
	cmd.Data_Dir = "./test"
	cmd.URL = testServer.URL

	message, err := cmd.ListFile()

	if err != nil {
		t.Error()
	}

	assert.Contains(t, message, `{"files": ["new.yaml","new1.txt"] }`)

}
