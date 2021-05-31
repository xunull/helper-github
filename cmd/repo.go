package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xunull/helper-github/pkg/console"
	"github.com/xunull/helper-github/pkg/option"
)

var (
	RepoType      string
	RepoLang      string
	RepoTopic     string
	RepoSort      string
	RepoAfterDate string
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "repo",
	Long:  `repo`,
	Run: func(cmd *cobra.Command, args []string) {
		console.ListRepos(option.WithRepoType(RepoType),
			option.WithRepoLang(RepoLang),
			option.WithOutputType(OutputType),
			option.WithUseCache(UseCache),
			option.WithAfterStr(RepoAfterDate),
			option.WithRepoSort(RepoSort))
	},
}

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	Run: func(cmd *cobra.Command, args []string) {
		console.ListRepos(option.WithRepoType(RepoType),
			option.WithRepoLang(RepoLang),
			option.WithOutputType(OutputType),
			option.WithUseCache(UseCache),
			option.WithAfterStr(RepoAfterDate),
			option.WithRepoSort(RepoSort))
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoListCmd)

	repoCmd.PersistentFlags().StringVar(&RepoType, "type", "", "repo type,private or public")
	repoCmd.PersistentFlags().StringVar(&RepoLang, "lang", "", "repo language, like go, python, java")
	repoCmd.PersistentFlags().StringVar(&RepoSort, "sort", "", "repo sort, create_time")
	repoCmd.PersistentFlags().StringVar(&RepoAfterDate, "after", "", "after")
}
