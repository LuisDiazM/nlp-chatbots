package config

import "github.com/spf13/viper"

type Env struct {
	PORT            int    `mapstructure:"PORT"`
	MONGO_URL       string `mapstructure:"MONGO_URL"`
	MONGO_USER      string `mapstructure:"MONGO_USER"`
	MONGO_PASSWORD  string `mapstructure:"MONGO_PASSWORD"`
	MONGO_PORT      string `mapstructure:"MONGO_PORT"`
	MONGO_POOL_SIZE uint64 `mapstructure:"MONGO_POOL_SIZE"`
	NATS_URL        string `mapstructure:"NATS_URL"`
	NATS_PORT       string `mapstructure:"NATS_PORT"`
	NATS_USER       string `mapstructure:"NATS_USER"`
	NATS_PASSWORD   string `mapstructure:"NATS_PASSWORD"`
}

func NewEnvironmentsSpecification() *Env {
	viper.AutomaticEnv()
	env := Env{PORT: viper.GetInt("PORT"),
		MONGO_URL:       viper.GetString("MONGO_URL"),
		MONGO_POOL_SIZE: viper.GetUint64("MONGO_POOL_SIZE"),
		NATS_URL:        viper.GetString("NATS_URL"),
		NATS_PORT:       viper.GetString("NATS_PORT"),
		NATS_USER:       viper.GetString("NATS_USER"),
		NATS_PASSWORD:   viper.GetString("NATS_PASSWORD"),
		MONGO_USER:      viper.GetString("MONGO_USER"),
		MONGO_PASSWORD:  viper.GetString("MONGO_PASSWORD"),
		MONGO_PORT:      viper.GetString("MONGO_PORT")}
	return &env
}
