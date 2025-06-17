package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8s-controller-tutorial",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
`,
	Run: func(cmd *cobra.Command, args []string) {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg("I warn you")
		log.Debug().Msg("Show me the code")
		log.Trace().Msg("I want to see everything")
		log.Warn().Msg("Not yet, but close")
		log.Error().Msg("If you see it, we have a problem")
		log.Fatal().Msg("I'm done")
		fmt.Println("Welcome to k8s-controller-tutorial CLI!")
	},
}

func Execute() {
	rootCmd.Execute()
}
