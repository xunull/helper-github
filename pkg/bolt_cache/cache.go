package bolt_cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/xunull/helper-github/internal/global"
	"github.com/xunull/helper-github/pkg/option"
	"github.com/xunull/helper-github/pkg/simple_api/github_api_v3/gh_model"
)

var (
	NoBoltError      = errors.New("no bolt db")
	NoTargetKeyError = errors.New("no target key")
)

func SaveUserRepos(repos []*gh_model.Repository, opts ...option.Option) {
	op := option.GetOption(opts...)
	username := op.TargetUser
	if op.TargetUser == "" {
		username = global.Config.User
	}

	b, err := json.Marshal(op)

	if err != nil {
		log.Error().Err(err).Msgf("SaveUserRepos Failed")
	} else {
		key := fmt.Sprintf("%s %s", username, string(b))
		SetCache(key, repos)
	}
}

func GetUserRepos(opts ...option.Option) ([]*gh_model.Repository, error) {
	op := option.GetOption(opts...)
	username := op.TargetUser
	if op.TargetUser == "" {
		username = global.Config.User
	}

	b, err := json.Marshal(op)

	if err == nil {
		key := fmt.Sprintf("%s %s", username, string(b))
		var repos []*gh_model.Repository
		err = GetCache(key, &repos)
		if err != nil {
			return nil, err
		} else {
			return repos, nil
		}
	} else {
		return nil, err
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func SaveUserStars(repos []*gh_model.Repository, opts ...option.Option) {
	op := option.GetOption(opts...)
	username := op.TargetUser
	if op.TargetUser == "" {
		username = global.Config.User
	}

	b, err := json.Marshal(op)

	if err != nil {
		log.Error().Err(err).Msgf("SaveUserRepos Failed")
	} else {
		key := fmt.Sprintf("%s stars %s", username, string(b))
		SetCache(key, repos)
	}
}

func GetUserStars(opts ...option.Option) ([]*gh_model.Repository, error) {
	op := option.GetOption(opts...)
	username := op.TargetUser
	if op.TargetUser == "" {
		username = global.Config.User
	}

	b, err := json.Marshal(op)

	if err == nil {
		key := fmt.Sprintf("%s stars %s", username, string(b))
		var repos []*gh_model.Repository
		err = GetCache(key, &repos)
		if err != nil {
			return nil, err
		} else {
			return repos, nil
		}
	} else {
		return nil, err
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func SetCache(key string, v interface{}) {

	b, err := json.Marshal(v)
	if err != nil {
		log.Debug().Err(err).Msgf("SetCache Failed: key is %s", key)
	} else {
		if global.BoltCache != nil {
			err := global.BoltCache.SetCache(key, b)
			if err != nil {
				log.Debug().Err(err).Msgf("SetCache Failed: key is %s", key)
			} else {
				log.Debug().Msgf("SetCache Success: key is %s", key)
			}
		} else {
			log.Debug().Msg("SetCache but BoltCache is nil")
		}
	}

}

func GetCache(key string, v interface{}) error {
	if global.BoltStore != nil {
		b, err := global.BoltCache.GetCache(key)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, v)
	} else {
		return NoBoltError
	}
}
