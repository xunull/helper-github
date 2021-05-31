package gh_api

import (
	"context"
	"github.com/google/go-github/v35/github"
	"github.com/xunull/helper-github/internal/global"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_log"
)

// ---------------------------------------------------------------------------------------------------------------------

func SearchRepoByTopic(topic string, opts ...option.Option) ([]*github.Repository, *github.Response, error) {

	q := "topic:" + topic

	op := option.GetOption(opts...)

	if op.RepoListLimit == 0 {
		op.RepoListLimit = global.Config.TopicListLimit
	}
	page := 1
	size := 100

	// todo, <100 may be not useful
	if global.Config.TopicListLimit > size {
		size = global.Config.TopicListLimit
	}

	var searchResult *github.RepositoriesSearchResult
	var repos []*github.Repository
	var resp *github.Response
	var err error

	searchOption := &github.SearchOptions{
		Sort: "stars",
	}
	searchOption.Page = page
	searchOption.PerPage = size

	for {

		searchResult, resp, err = client.Search.Repositories(context.Background(), q, searchOption)

		if err != nil {
			simple_log.LogError(err)
			return nil, resp, err
		}

		repos = append(repos, searchResult.Repositories...)

		if len(repos) > op.RepoListLimit {
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
