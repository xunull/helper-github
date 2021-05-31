package gh_api

import (
	"github.com/xunull/goc/commonx"
	"testing"
)

func TestGetUserAllStarred(t *testing.T) {
	initTestClient(t)
	repos, resp, err := GetUserAllStarred("xunull")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", commonx.JsonString(repos))
		t.Logf("%+v", resp)
	}
}

func TestGetUserStarred(t *testing.T) {
	initTestClient(t)
	repos, resp, err := GetUserStarred("xunull", 1, 50)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", commonx.JsonString(repos))
		t.Logf("%+v", resp)
	}
}

func TestGetPlainUserStarred(t *testing.T) {
	initTestClient(t)
	repos, err := GetPlainUserStarred("xunull", 1, 50)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", commonx.JsonString(repos))
	}
}
