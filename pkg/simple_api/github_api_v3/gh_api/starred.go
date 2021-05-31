package gh_api

import (
	"context"
	"github.com/google/go-github/v35/github"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_model"
	"github.com/xunull/helper-github/pkg/simple_log"
)

func GetPlainUserStarred(username string, page, size int) ([]*gh_model.Repository, error) {
	repos, _, err := GetUserStarred(username, page, size)
	if err != nil {
		return nil, err
	}
	sim := make([]*gh_model.Repository, 0, len(repos))
	for _, r := range repos {
		s := &gh_model.Repository{}
		err = copier.Copy(s, r.Repository)
		sim = append(sim, s)
	}
	return sim, err
}

func GetPlainUserAllStarred(username string, opts ...option.Option) ([]*gh_model.Repository, error) {
	repos, _, err := GetUserAllStarred(username, opts...)
	if err != nil {
		return nil, err
	}
	sim := make([]*gh_model.Repository, 0, len(repos))
	for _, r := range repos {
		s := &gh_model.Repository{}
		err = copier.Copy(s, r.Repository)
		sim = append(sim, s)
	}
	return sim, err
}

func GetUserStarred(username string, page, size int) ([]*github.StarredRepository, *github.Response, error) {
	r, resp, err := client.Activity.ListStarred(context.Background(), username, &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: size,
		},
	})
	if err != nil {
		simple_log.LogError(err)
		return r, resp, err
	}
	return r, resp, err
}

func GetUserAllStarred(username string, opts ...option.Option) ([]*github.StarredRepository, *github.Response, error) {
	size := 100
	page := 1
	var repos, r []*github.StarredRepository
	var resp *github.Response
	var err error
	for {
		r, resp, err = client.Activity.ListStarred(context.Background(), username, &github.ActivityListStarredOptions{
			ListOptions: github.ListOptions{
				Page:    page,
				PerPage: size,
			},
		})
		if err != nil {
			simple_log.LogError(err)
			return repos, resp, err
		}
		repos = append(repos, r...)
		if resp.NextPage > 0 {
			page++
		} else {
			break
		}
		log.Debug().Msgf("%+v", resp)
	}
	return repos, resp, err
}
