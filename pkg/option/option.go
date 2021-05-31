package option

import (
	"github.com/xunull/goc/enhance/timex"
	"time"
)

type option struct {
	RepoType      string
	RepoLang      string
	RepoTopic     string
	RepoSort      string
	RepoAfter     time.Time
	RepoListLimit int
	OutputType    string
	TargetUser    string
	UseCache      bool `json:"-"`
	TopicOpenUrl  bool
}

type Option func(*option)

func GetOption(opts ...Option) *option {
	op := &option{}
	for _, opt := range opts {
		opt(op)
	}
	return op
}

func WithTopicOpenUrl(open bool) Option {
	return func(o *option) {
		o.TopicOpenUrl = open
	}
}

func WithAfterStr(after string) Option {
	return func(o *option) {
		a, err := timex.GetYYYYMMDDTime(after)
		if err == nil {
			o.RepoAfter = a
		} else {
			a, err := timex.GetYYYYMMTime(after)
			if err == nil {
				o.RepoAfter = a
			} else {
				a, err := timex.GetYYYYTime(after)
				if err == nil {
					o.RepoAfter = a
				}
			}
		}
	}
}

func WithRepoListLimit(limit int) Option {
	return func(o *option) {
		o.RepoListLimit = limit
	}
}

func WithRepoSort(name string) Option {
	return func(o *option) {
		o.RepoSort = name
	}
}

func WithRepoTopic(topic string) Option {
	return func(o *option) {
		o.RepoTopic = topic
	}
}

func WithUseCache(use bool) Option {
	return func(o *option) {
		o.UseCache = use
	}
}

func WithTargetUser(user string) Option {
	return func(o *option) {
		o.TargetUser = user
	}
}

func WithRepoLang(lang string) Option {
	return func(o *option) {
		o.RepoLang = lang
	}
}

func WithOutputType(t string) Option {
	return func(o *option) {
		o.OutputType = t
	}
}

func WithRepoType(t string) Option {
	return func(o *option) {
		o.RepoType = t
	}
}
