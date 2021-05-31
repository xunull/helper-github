package console

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"github.com/xunull/goc/commonx"
	"github.com/xunull/goc/enhance/timex"
	"github.com/xunull/helper-github/pkg/bolt_cache"
	"github.com/xunull/helper-github/pkg/config"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_api"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_model"
	"strconv"
	"strings"
)

func ListRepos(opts ...option.Option) {

	op := option.GetOption(opts...)

	var repos []*gh_model.Repository
	var err error
	if op.UseCache {
		repos, err = bolt_cache.GetUserRepos(opts...)
		if err != nil {
			log.Trace().Err(err).Msgf("GetUserRepos Failed")
		}
	}

	if repos == nil {

		if op.RepoAfter.IsZero() {
			repos, err = gh_api.GetPlainUserAllRepos(op.TargetUser, opts...)
		} else {
			repos, err = gh_api.GetPlainUserAllRepoAfterDate(op.TargetUser, op.RepoAfter, opts...)
		}

		commonx.CheckErrOrFatal(err)
		bolt_cache.SaveUserRepos(repos, opts...)
	}

	gh_model.SortRepoListByLang(repos)

	if op.RepoLang != "" {
		op.RepoLang = strings.ToLower(op.RepoLang)
		res := gh_model.GroupRepoListByLang(repos)
		if _, ok := res[op.RepoLang]; ok {
			repos = res[op.RepoLang]
		} else {
			fmt.Println("target language not find")
		}
	}

	if op.RepoSort != "" {
		if op.RepoSort == "create_time" {
			var r gh_model.RepoList
			r = repos
			r.SortRepoListByCreateTime()
		}
	} else {
		if !op.RepoAfter.IsZero() {
			var r gh_model.RepoList
			r = repos
			r.SortRepoListByCreateTime()
		}
	}

	OutputRepos(repos, opts...)
}

// ---------------------------------------------------------------------------------------------------------------------

func OutputRepos(repos []*gh_model.Repository, opts ...option.Option) {
	op := option.GetOption(opts...)
	if op.OutputType == config.OutputTypeJson {
		OutputReposJson(repos)
	} else if op.OutputType == config.OutputTypeUi {
		OutputReposUi(repos)
	} else {
		OutputReposTable(repos)
	}
}

func OutputReposJson(repos []*gh_model.Repository) {
	str := commonx.JsonString(repos)
	fmt.Println(str)
}

func OutputReposTable(repos []*gh_model.Repository) {
	tableStr := &strings.Builder{}
	table := tablewriter.NewWriter(tableStr)
	table.SetHeader([]string{"Index", "Id", "Name", "Lang", "CreateAt"})

	for i, repo := range repos {
		data := []string{
			strconv.Itoa(i + 1),
			strconv.Itoa(int(repo.ID)),
			repo.Name,
			repo.Language,
			timex.GetTimeYYYYMMDD(repo.CreatedAt.Time),
		}
		table.Append(data)
	}
	table.Render()
	fmt.Println(tableStr.String())
}

func OutputReposUi(repos []*gh_model.Repository, opts ...option.Option) {
	op := option.GetOption(opts...)
	var showLang bool
	if op.RepoLang == "" {
		showLang = true
	}

	app := tview.NewApplication()
	table := tview.NewTable()

	for row, repo := range repos {
		idCell := tview.NewTableCell(strconv.Itoa(int(repo.ID)))
		table.SetCell(row, 0, idCell)
		nameCell := tview.NewTableCell(repo.Name)
		table.SetCell(row, 1, nameCell)
		if showLang {
			langCell := tview.NewTableCell(repo.Language)
			table.SetCell(row, 2, langCell)
		}
	}

	table.SetBorders(true).SetTitle("Repos")
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
