package global

import (
	"github.com/xunull/goc/boltkv"
	"github.com/xunull/goc/boltkv_cache"
	"github.com/xunull/goc/myhome"
)

var (
	AppName      = "helper-github"
	Debug        bool
	Initializing bool
	BoltStore    *boltkv.BoltStore
	BoltCache    *boltkv_cache.CacheStore
)

var (
	MyHome *myhome.MyHome
)

func CloseBolt() {
	if BoltStore != nil {
		BoltStore.Close()
	}
}
