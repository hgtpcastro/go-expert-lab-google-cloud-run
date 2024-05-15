package main

import (
	"os"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/services/weather-service/internal/shared/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "zip-code-microservice",
	Short:            "zip code microservice",
	Long:             "zip code microservice",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		app.NewApp().Run()
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
