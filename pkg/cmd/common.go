package cmd

import (
	"context"
	"space-traders-playground/pkg/repository/mongo"
	"space-traders-playground/pkg/service"

	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Email    string `envconfig:"PLAYER_EMAIL"`
	Callsign string `envconfig:"PLAYER_CALLSIGN"`
}

func NewGameService(ctx context.Context, logger *zap.Logger) service.Service {

	var conf mongo.Config

	if err := envconfig.Init(&conf); err != nil {
		logger.Fatal(err.Error())
	}

	repo, err := mongo.NewRepository(conf, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	svc, err := service.NewService(logger, repo)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return svc

}
