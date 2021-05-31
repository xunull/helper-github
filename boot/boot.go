package boot

import (
	"github.com/rs/zerolog/log"
	"github.com/xunull/goc/boltkv"
	"github.com/xunull/goc/boltkv_cache"
	"github.com/xunull/helper-github/internal/global"
	"time"
)

func Boot() {
	global.MyHome.Init()
	initBoltCache()
}

func initBoltCache() {
	dbName := "hg-bolt"
	db, err := boltkv.InitBoltStore(dbName, global.MyHome.GetDbHome())
	if err != nil {
		log.Error().Err(err).Msgf("initBoltCache Failed")
	} else {
		global.BoltStore = db
		cache := boltkv_cache.NewCacheStore(*db,
			boltkv_cache.WithDuration(time.Duration(global.Config.CacheDuration)*time.Second))
		global.BoltCache = cache
	}
}
