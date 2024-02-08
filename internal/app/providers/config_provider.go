package providers

import (
	"fmt"
	"go-assignment-bootcamp/internal/app/helpers"
	"go-assignment-bootcamp/internal/config"

	"github.com/spf13/viper"
)

type ConfigProvider struct {
	config *helpers.Config
}

func NewConfigProvider() *ConfigProvider {
	configProvider := &ConfigProvider{
		config: helpers.NewConfig(),
	}
	configProvider.register()
	return configProvider
}

func (configProvider *ConfigProvider) GetConfig() *helpers.Config {
	return configProvider.config
}

func (configProvider *ConfigProvider) register() {
	configProvider.initViper()

	configProvider.config.Add(config.GetApplicationConfigs())
	configProvider.config.Add(config.GetDatabaseConfigs())
	configProvider.config.Add(config.GetHTTPConfigs())
}

func (configProvider *ConfigProvider) initViper() {
	viper.SetConfigType("env")
	viper.SetConfigFile("../../.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
