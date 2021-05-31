package gh_api

import (
	"context"
	"github.com/google/go-github/v35/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var client *github.Client
var logger *zap.Logger

func initLogger() {
	logger, _ = zap.NewProduction()
}

func InitClient(token string) *github.Client {

	initLogger()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := github.NewClient(tc)

	client = c

	return c
}
