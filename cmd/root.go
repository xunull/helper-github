package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xunull/helper-github/boot"
	"github.com/xunull/helper-github/internal/global"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_api"
	"log"
	"os"
	"runtime/debug"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	Debug       bool
	TargetUser  string
	TargetToken string
	OutputType  string
	UseCache    bool
)

var rootCmd = &cobra.Command{
	Use:   "helper-github",
	Short: "helper-github",
	Long:  `helper-github`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		gh_api.InitClient(global.Config.Token)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		global.CloseBolt()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.helper-github.yaml)")
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "debug")
	rootCmd.PersistentFlags().StringVarP(&TargetUser, "user", "u", "", "target user")
	rootCmd.PersistentFlags().StringVarP(&TargetToken, "token", "t", "", "target token")

	rootCmd.PersistentFlags().StringVar(&OutputType, "out", "table", "output type")
	rootCmd.PersistentFlags().BoolVar(&UseCache, "cache", false, "use cache")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".helper-github")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if Debug {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
		var config global.AppConfig
		err := viper.Unmarshal(&config)
		if err != nil {
			debug.PrintStack()
			log.Fatal(err)
		}
		global.InitConfig(&config)
		boot.Boot()
	}
}
