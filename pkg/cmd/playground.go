package cmd

import (
	"context"

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

	var conf Config

	if err := envconfig.Init(&conf); err != nil {
		logger.Fatal(err.Error())
	}
	svc := NewGameService(ctx, logger)
	svc.RegisterNewAgent(ctx, "UNITED", conf.Email, conf.Callsign)
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
