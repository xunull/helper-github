package gh_api

import (
	"github.com/xunull/goc/commonx"
	"testing"
)

func TestGetPlainUserRepo(t *testing.T) {
	initTestClient(t)
	repos, err := GetPlainUserRepo("xunull", 1, 50)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", commonx.JsonString(repos))
	}
}
