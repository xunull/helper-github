package console

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/xunull/goc/commonx"
	"github.com/xunull/helper-github/pkg/bolt_cache"
	"github.com/xunull/helper-github/pkg/config"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_api"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_model"
)

func ListStars(opts ...option.Option) {
	op := option.GetOption(opts...)

	var repos []*gh_model.Repository
	var err error
	if op.UseCache {
		repos, err = bolt_cache.GetUserStars(opts...)
		if err != nil {
			log.Trace().Err(err).Msgf("GetUserRepos Failed")
		}
	}

	if repos == nil {
		repos, err = gh_api.GetPlainUserAllStarred(op.TargetUser, opts...)
		commonx.CheckErrOrFatal(err)
		bolt_cache.SaveUserStars(repos, opts...)
	}

	commonx.CheckErrOrFatal(err)

	gh_model.SortRepoListByLang(repos)

	if op.RepoLang != "" {
		res := gh_model.GroupRepoListByLang(repos)
		if _, ok := res[op.RepoLang]; ok {
			repos = res[op.RepoLang]
		} else {
			fmt.Println("target language not find")
		}
	}

	if op.RepoTopic != "" {
		var r gh_model.RepoList
		r = repos
		topicMap := r.GetRepoWithTopic(op.RepoTopic)
		if _, ok := topicMap[op.RepoTopic]; !ok {
			fmt.Println("no target topic")
			return
		}
		repos = topicMap[op.RepoTopic]
	}

	if op.OutputType == config.OutputTypeJson {
		OutputReposJson(repos)
	} else if op.OutputType == config.OutputTypeUi {
		OutputReposUi(repos)
	} else {
		OutputReposTable(repos)
	}
}
