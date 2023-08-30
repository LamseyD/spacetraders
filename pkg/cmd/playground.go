package cmd

import (
	"github.com/spf13/cobra"
)

func play(args []string) {

}

var playgroundCmd = &cobra.Command{
    Use:   "play",
    Short: "play SpaceTraders",
    Run:   func(cmd *cobra.Command, args []string) {
		play(args)
	},
}

func init() {
    rootCmd.AddCommand(playgroundCmd)
}
