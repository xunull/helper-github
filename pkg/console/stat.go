package console

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/xunull/goc/commonx"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_api"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_model"
	"log"
)

func StatLang(opts ...option.Option) {
	op := option.GetOption(opts...)
	repos, err := gh_api.GetPlainUserAllRepos(op.TargetUser, opts...)
	commonx.CheckErrOrFatal(err)

	res := gh_model.StatRepoListByLang(repos)

	OutputStatLang(res)
}

func OutputStatLang(data map[string]int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	ui.Clear()
	defer ui.Close()

	labels := make([]string, 0, len(data))
	bcData := make([]float64, 0, len(data))
	for key, value := range data {
		labels = append(labels, key)
		bcData = append(bcData, float64(value))
	}
	bc := widgets.NewBarChart()
	bc.Labels = labels
	bc.Data = bcData
	bc.Title = "Lang"
	bc.SetRect(5, 5, 100, 30)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func StatCreateTime(opts ...option.Option) {
	op := option.GetOption(opts...)
	repos, err := gh_api.GetPlainUserAllRepos(op.TargetUser, opts...)
	commonx.CheckErrOrFatal(err)

	var r gh_model.RepoList
	r = repos
	res := r.StatRepoListByByCreateTime()

	OutputStat("year", res)
}

func OutputStat(title string, data map[string]int) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	ui.Clear()
	defer ui.Close()

	labels := make([]string, 0, len(data))
	bcData := make([]float64, 0, len(data))
	for key, value := range data {
		labels = append(labels, key)
		bcData = append(bcData, float64(value))
	}
	bc := widgets.NewBarChart()
	bc.Labels = labels
	bc.Data = bcData
	bc.Title = title
	bc.SetRect(5, 5, 100, 30)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
