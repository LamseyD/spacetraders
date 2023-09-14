package cmd

import (
	"context"
	"space-traders-playground/pkg/repository/mongo"
	"space-traders-playground/pkg/service"

	"github.com/spf13/cobra"
	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
)

func play(args []string) {
	ctx := context.Background()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	var conf mongo.Config

	if err := envconfig.Init(&conf); err != nil {
		logger.Fatal(err.Error())
	}

	repo, err := mongo.NewRepository(conf, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	svc := service.NewService(logger, repo)
	svc.RegisterNewAgent(ctx, "test", "test", "test")
}

var playgroundCmd = &cobra.Command{
	Use:   "play",
	Short: "play SpaceTraders",
	Run: func(cmd *cobra.Command, args []string) {
		play(args)
	},
}

func init() {
	rootCmd.AddCommand(playgroundCmd)
}
