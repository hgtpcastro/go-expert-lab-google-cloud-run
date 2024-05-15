package environment

import (
	"log"
	"os"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/constants"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Environment string

var (
	Development = Environment(constants.Development)
	Production  = Environment(constants.Production)
	Test        = Environment(constants.Test)
)

func ProvideAppEnv(environments ...Environment) Environment {
	environment := Environment("")

	if len(environments) > 0 {
		environment = environments[0]
	} else {
		environment = Development
	}

	// setup viper to read from os environment with `viper.Get`
	viper.AutomaticEnv()

	// https://articles.wesionary.team/environment-variable-configuration-in-your-golang-project-using-viper-4e8289ef664d
	// load environment variables form .env files to system environment variables,
	// it just find `.env` file in our current `executing working directory` in our app for example `zip_code_service`
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file cannot be found.")
	}

	manualEnv := os.Getenv(constants.AppEnv)

	if manualEnv != "" {
		environment = Environment(manualEnv)
	}

	return environment
}

func (e Environment) IsDevelopment() bool {
	return e == Development
}

func (e Environment) IsProduction() bool {
	return e == Production
}

func (e Environment) IsTest() bool {
	return e == Test
}
