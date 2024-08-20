package configs

import (
	"strings"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/weather_service/pkg/configs"
)

type AppOptions struct {
	ServiceName string `mapstructure:"serviceName"`
}

type ConfigOptions struct {
	AppOptions AppOptions `mapstructure:"appOptions"`
}

func ProvideConfig(environment string) (*ConfigOptions, error) {
	return configs.BindConfigKey[*ConfigOptions]("", environment)
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}
