package gh_api

import (
	"context"
	"github.com/google/go-github/v35/github"
	"github.com/jinzhu/copier"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_model"
)

func GetUserInfo(username string) (*github.User, *github.Response, error) {
	user, resp, err := client.Users.Get(context.Background(), username)
	if err != nil {
		logger.Sugar().Errorf("GetUserInfo Failed: %s", err)
	}
	return user, resp, err
}

func GetSimpleUserInfo(username string) (*gh_model.User, error) {
	user, _, err := GetUserInfo(username)
	if err != nil {
		return nil, err
	}
	s := &gh_model.User{}
	err = copier.Copy(s, user)
	return s, nil
}
