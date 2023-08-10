package cmd_test

import (
	"net/http"
	"net/http/httptest"
	"store-cli/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWC(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"items": [
			  {
				"Word": "rm",
				"Frequency": 1
			  },
			  {
				"Word": "on",
				"Frequency": 1
			  },
			  {
				"Word": "controller:",
				"Frequency": 1
			  },
			  {
				"Word": "shell",
				"Frequency": 1
			  },
			  {
				"Word": "command:",
				"Frequency": 1
			  },
			  {
				"Word": "same",
				"Frequency": 1
			  },
			  {
				"Word": "parallscssc",
				"Frequency": 1
			  },
			  {
				"Word": "absent.",
				"Frequency": 1
			  },
			  {
				"Word": "metadata:",
				"Frequency": 1
			  },
			  {
				"Word": "List",
				"Frequency": 1
			  }
			],
			"totalWordCount": 192
	   }`))
	}))

	cmd.URL = testServer.URL

	message, err := cmd.WC("10", "a")

	if err != nil {
		t.Error()
	}

	assert.Contains(t, message, "totalWordCount")

}
