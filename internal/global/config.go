package global

import (
	"github.com/mitchellh/go-homedir"
	"github.com/xunull/goc/commonx"
	"github.com/xunull/goc/myhome"
	"path/filepath"
)

const (
	DefaultTopicListLimit = 50
	TopicOpenUrlMax       = 10
)

var Config *AppConfig

type AppConfig struct {
	Home            string
	User            string
	Token           string
	CacheDuration   int
	TopicListLimit  int
	TopicOpenUrlMax int
}

func fillConfig(config *AppConfig) {
	if config.Home == "" {
		config.Home = GetHome()
	}
	home, err := homedir.Expand(config.Home)
	commonx.CheckErrOrFatal(err)
	config.Home = home

	MyHome = myhome.NewMyHome(home)

	if config.TopicListLimit == 0 {
		config.TopicListLimit = DefaultTopicListLimit
	}

	if config.TopicOpenUrlMax == 0 {
		config.TopicOpenUrlMax = TopicOpenUrlMax
	}

}

func InitConfig(config *AppConfig) {
	fillConfig(config)
	Config = config
}

func GetHome() string {
	home, err := homedir.Dir()
	commonx.CheckErrOrFatal(err)
	home = filepath.Join(home, ".helper-github")
	return home
}
