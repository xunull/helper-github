package gh_api

import (
	"context"
	"github.com/google/go-github/v35/github"
	"github.com/jinzhu/copier"
	"github.com/xunull/helper-github/pkg/config"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_model"
	"github.com/xunull/helper-github/pkg/simple_log"
	"time"
)

func GetPlainUserRepo(username string, page, size int) ([]*gh_model.Repository, error) {
	repos, _, err := GetUserRepo(username, page, size)
	if err != nil {
		return nil, err
	}
	sim := make([]*gh_model.Repository, 0, len(repos))
	for _, r := range repos {
		s := &gh_model.Repository{}
		err = copier.Copy(s, r)
		sim = append(sim, s)
	}
	return sim, err
}

func GetPlainUserAllRepos(username string, opts ...option.Option) ([]*gh_model.Repository, error) {
	repos, _, err := GetUserAllRepos(username, opts...)
	if err != nil {
		return nil, err
	}
	sim := make([]*gh_model.Repository, 0, len(repos))
	for _, r := range repos {
		s := &gh_model.Repository{}
		err = copier.Copy(s, r)
		sim = append(sim, s)
	}
	return sim, err
}

func GetPlainUserAllRepoAfterDate(username string, date time.Time, opts ...option.Option) ([]*gh_model.Repository, error) {
	repos, _, err := GetUserAllRepoAfterDate(username, date, opts...)
	if err != nil {
		return nil, err
	}
	sim := make([]*gh_model.Repository, 0, len(repos))
	for _, r := range repos {
		s := &gh_model.Repository{}
		err = copier.Copy(s, r)
		sim = append(sim, s)
	}
	return sim, err
}

// ---------------------------------------------------------------------------------------------------------------------

func GetUserRepo(username string, page, size int) ([]*github.Repository, *github.Response, error) {
	r, resp, err := client.Repositories.List(context.Background(), username, &github.RepositoryListOptions{
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

func GetUserAllRepoAfterDate(username string, date time.Time, opts ...option.Option) ([]*github.Repository, *github.Response, error) {
	op := option.GetOption(opts...)
	size := 100
	page := 1
	var repos, r []*github.Repository
	var resp *github.Response
	var err error
	repoListOption := &github.RepositoryListOptions{
		Sort: "created",
	}
	if op.RepoType == config.RepoTypePrivate {
		repoListOption.Visibility = config.RepoTypePrivate
	} else if op.RepoType == config.RepoTypePublic {
		repoListOption.Visibility = config.RepoTypePublic
	} else {
		repoListOption.Visibility = config.RepoTypeAll
	}
	for {
		repoListOption.ListOptions = github.ListOptions{
			Page:    page,
			PerPage: size,
		}
		r, resp, err = client.Repositories.List(context.Background(), username, repoListOption)
		if err != nil {
			simple_log.LogError(err)
			return repos, resp, err
		}
		flag := false
		for _, repo := range r {
			if repo.CreatedAt.After(date) {
				repos = append(repos, repo)
			} else {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		if resp.NextPage > 0 {
			page++
		} else {
			break
		}
	}
	return repos, resp, err
}

func GetUserAllRepos(username string, opts ...option.Option) ([]*github.Repository, *github.Response, error) {
	op := option.GetOption(opts...)
	size := 100
	page := 1
	var repos, r []*github.Repository
	var resp *github.Response
	var err error

	repoListOption := &github.RepositoryListOptions{

	}
	if op.RepoType == config.RepoTypePrivate {
		repoListOption.Visibility = config.RepoTypePrivate
	} else if op.RepoType == config.RepoTypePublic {
		repoListOption.Visibility = config.RepoTypePublic
	} else {
		repoListOption.Visibility = config.RepoTypeAll
	}

	for {
		repoListOption.ListOptions = github.ListOptions{
			Page:    page,
			PerPage: size,
		}
		r, resp, err = client.Repositories.List(context.Background(), username, repoListOption)

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
	}
	return repos, resp, err
}
