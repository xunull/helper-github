package cmd

import (
	"github.com/xunull/helper-github/pkg/console"

	"github.com/spf13/cobra"
)

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "stat",
	Long:  `stat`,
}

var statLangCmd = &cobra.Command{
	Use:   "lang",
	Short: "lang",
	Long:  `lang`,
	Run: func(cmd *cobra.Command, args []string) {
		console.StatLang()
	},
}

var statYearTimeCmd = &cobra.Command{
	Use:   "year",
	Short: "year",
	Long:  `year`,
	Run: func(cmd *cobra.Command, args []string) {
		console.StatCreateTime()
	},
}

func init() {
	rootCmd.AddCommand(statCmd)
	statCmd.AddCommand(statLangCmd)
	statCmd.AddCommand(statYearTimeCmd)
}
