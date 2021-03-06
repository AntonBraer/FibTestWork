package config

import "github.com/spf13/viper"

type Config struct {
	MainRedis string `mapstructure:"MAIN_REDIS_ADDRESS"`
	HttpUrl   string `mapstructure:"HTTP_URL"`
	GrpcPort  string `mapstructure:"GRPC_PORT"`
}

// LoadConfig загружает конфиг из app.env
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
