package simple_log

import "github.com/rs/zerolog/log"

func LogError(err error) {
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}
