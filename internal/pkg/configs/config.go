package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v8"
	typemapper "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/reflection/type_mapper"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

func BindConfig[T any](environment string) (T, error) {
	return BindConfigKey[T]("", environment)
}

func BindConfigKey[T any](configKey string, environment string) (T, error) {
	var configPath string

	// https://articles.wesionary.team/environment-variable-configuration-in-your-golang-project-using-viper-4e8289ef664d
	configPathFromEnv := viper.Get("CONFIG_PATH")
	if configPathFromEnv != nil {
		configPath = configPathFromEnv.(string)
	} else {
		// https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
		// https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
		d, err := getConfigRootPath()
		if err != nil {
			return *new(T), err
		}

		configPath = d
	}

	cfg := typemapper.GenericInstanceByT[T]()

	// this should set before reading config values from json file
	// https://github.com/mcuadros/go-defaults
	defaults.SetDefaults(cfg)

	// https://github.com/spf13/viper/issues/390#issuecomment-718756752
	viper.SetConfigName(fmt.Sprintf("config.%s", environment))
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return *new(T), err
	}

	if len(configKey) == 0 {
		if err := viper.Unmarshal(cfg); err != nil {
			return *new(T), err
		}
	} else {
		if err := viper.UnmarshalKey(configKey, cfg); err != nil {
			return *new(T), err
		}
	}

	viper.AutomaticEnv()

	// https://github.com/caarlos0/env
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return cfg, nil
}

func getConfigRootPath() (string, error) {
	// Get the current working directory
	// Getwd gives us the current working directory that we are running our app with `go run` command.
	// In goland we can specify this working directory for the project
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fmt.Printf("Current working directory is: %s\n", wd)

	// Get the absolute path of the executed project directory
	absCurrentDir, err := filepath.Abs(wd)
	if err != nil {
		return "", err
	}

	// Get the path to the "config" folder within the project directory
	configPath := filepath.Join(absCurrentDir, "configs")

	return configPath, nil
}
