package gh_api

import (
	"github.com/xunull/goc/commonx"
	"os"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	token := os.Getenv("github_token")
	if token == "" {
		t.Error("No token")
	}
	InitClient(token)
	user, resp, err := GetUserInfo("xunull")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", commonx.JsonString(user))
		t.Logf("%v", resp)
	}
}
