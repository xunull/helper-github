package gh_api

import (
	"os"
	"testing"
)

func initTestClient(t *testing.T) {
	token := os.Getenv("github_token")
	if token == "" {
		t.Error("No token")
	}
	InitClient(token)
}
