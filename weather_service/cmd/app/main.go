package main

import (
	"os"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/weather_service/weather/shared/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "weather-microservice",
	Short:            "weather microservice",
	Long:             "weather microservice",
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
