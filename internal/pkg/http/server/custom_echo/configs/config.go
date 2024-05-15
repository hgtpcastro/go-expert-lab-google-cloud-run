package configs

import (
	"fmt"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/configs"
	typemapper "github.com/hgtpcastro/go-expert-lab-google-cloud-run/internal/pkg/reflection/type_mapper"
	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typemapper.GetTypeNameByT[EchoOptions]())

type EchoOptions struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

func (c *EchoOptions) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

func ProvideConfig(environment string) (*EchoOptions, error) {
	return configs.BindConfigKey[*EchoOptions](optionName, environment)
}
