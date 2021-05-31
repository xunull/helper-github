package console

import (
	"fmt"
	"github.com/google/go-github/v35/github"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
	"github.com/xunull/goc/enhance/timex"
	"github.com/xunull/helper-github/internal/global"
	"github.com/xunull/helper-github/pkg/config"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_api"
	"github.com/xunull/openurl"
	"strconv"
	"strings"
)

func SearchTopic(topic string, opts ...option.Option) {

	repos, _, err := gh_api.SearchRepoByTopic(topic, opts...)
	if err != nil {
		log.Error().Err(err).Msgf("SearchTopic Failed")
		return
	}

	OutputSearchTopic(repos, opts...)
	openTopicRepoUrl(repos, opts...)
}

func openTopicRepoUrl(repos []*github.Repository, opts ...option.Option) {
	op := option.GetOption(opts...)
	if op.TopicOpenUrl {
		urls := make([]string, 0)

		for _, repo := range repos {
			urls = append(urls, *repo.HTMLURL)
		}
		min := global.Config.TopicOpenUrlMax
		if min > len(urls) {
			min = len(urls)
		}
		for i := 0; i < min; i++ {
			openurl.Open(urls[i])
		}
	}
}

func OutputSearchTopic(repos []*github.Repository, opts ...option.Option) {
	op := option.GetOption(opts...)
	if op.OutputType == config.OutputTypeJson {
		OutputSearchTopicTable(repos)
	} else if op.OutputType == config.OutputTypeUi {
		OutputSearchTopicTable(repos)
	} else {
		OutputSearchTopicTable(repos)
	}
}

func OutputSearchTopicTable(repos []*github.Repository) {
	tableStr := &strings.Builder{}
	table := tablewriter.NewWriter(tableStr)
	table.SetHeader([]string{"Index", "Name", "Stars", "CreateAt"})

	for i, repo := range repos {
		data := []string{
			strconv.Itoa(i + 1),
			*repo.FullName,
			strconv.Itoa(*repo.StargazersCount),
			timex.GetTimeYYYYMMDD(repo.CreatedAt.Time),
		}
		table.Append(data)
	}
	table.Render()
	fmt.Println(tableStr.String())
}
