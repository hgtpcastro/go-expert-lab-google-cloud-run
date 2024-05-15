package configs

import (
	"strings"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/configs"
)

type AppOptions struct {
	ServiceName string `mapstructure:"serviceName"`
}

type ConfigOptions struct {
	AppOptions AppOptions `mapstructure:"appOptions"`
}

// func NewConfigOptions(environment string) (*ConfigOptions, error) {
// 	cfg, err := configs.BindConfig[*ConfigOptions](environment)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cfg, nil
// }

func ProvideConfig(environment string) (*ConfigOptions, error) {
	return configs.BindConfigKey[*ConfigOptions]("", environment)
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}
