package config

import "github.com/spf13/viper"

type ContextKey string
type Config struct {
	DBDriver                string `mapstructure:"DB_DRIVER"`
	DBSource                string `mapstructure:"DB_SOURCE"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	JwtSigningKey           string `mapstructure:"JWT_SINGNING_KEY"`
	JwtIssuer               string `mapstructure:"JWT_ISSUER"`
	JwtAccessTokenDuration  int    `mapstructure:"JWT_ACCESS_TOKEN_DURATION_SECONDS"`
	JwtRefreshTokenDuration int    `mapstructure:"JWT_REFRESH_TOKEN_DURATION_SECONDS"`
}

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
