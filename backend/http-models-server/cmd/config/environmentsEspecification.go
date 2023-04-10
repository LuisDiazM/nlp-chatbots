package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Port            int    `mapstructure:"PORT"`
	MONGO_URL       string `mapstructure:"MONGO_URL"`
	MONGO_POOL_SIZE int    `mapstructure:"MONGO_POOL_SIZE"`
}

func NewEnvironmentsSpecification() *Env {
	var env Env
	viper.SetConfigFile("cmd/config/.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &env
}
