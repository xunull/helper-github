package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xunull/helper-github/pkg/console"
	"github.com/xunull/helper-github/pkg/option"
)

var (
	TopicOpen bool
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "topic",
	Long:  `topic`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		console.SearchTopic(topic,
			option.WithRepoType(RepoType),
			option.WithRepoLang(RepoLang),
			option.WithRepoSort(RepoSort),
			option.WithOutputType(OutputType),
			option.WithUseCache(UseCache),
			option.WithTopicOpenUrl(TopicOpen))
	},
}

func init() {
	rootCmd.AddCommand(topicCmd)
	// open search result in browser
	topicCmd.PersistentFlags().BoolVar(&TopicOpen, "open", false, "open url, open top 10")
}
