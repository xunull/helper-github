package cmd

import (
	"github.com/xunull/helper-github/pkg/console"
	"github.com/xunull/helper-github/pkg/option"

	"github.com/spf13/cobra"
)

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "star",
	Long:  `star`,
	Run: func(cmd *cobra.Command, args []string) {
		console.ListStars(option.WithRepoLang(RepoLang),
			option.WithRepoTopic(RepoTopic),
			option.WithOutputType(OutputType),
			option.WithUseCache(UseCache),
		)
	},
}

var starListCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	Run: func(cmd *cobra.Command, args []string) {
		console.ListStars(option.WithRepoLang(RepoLang),
			option.WithRepoTopic(RepoTopic),
			option.WithOutputType(OutputType),
			option.WithUseCache(UseCache),
		)
	},
}

func init() {
	rootCmd.AddCommand(starCmd)
	starCmd.AddCommand(starListCmd)
	starCmd.PersistentFlags().StringVar(&RepoLang, "lang", "", "repo language, like go,python")
	starCmd.PersistentFlags().StringVar(&RepoTopic, "topic", "", "repo topic")
}
