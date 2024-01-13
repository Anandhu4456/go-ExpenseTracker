package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST     string `json:"db_host"`
	DB_USER 	string `json:"db_user"`
	DB_NAME     string `json:"db_name"`
	DB_PORT     string `json:"db_port"`
	DB_PASSWORD string `json:"db_password"`
}

var envs = []string{
	"DB_HOST","DB_USER", "DB_NAME", "DB_PORT", "DB_PASSWORD",
}

func LoadConfig() (Config, error) {
	var conf Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	for _, env := range envs {
		err := viper.BindEnv(env)
		if err != nil {
			return conf, err
		}
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading the env file")
	}
	if err = viper.Unmarshal(&conf); err != nil {
		return conf, err
	}
	return conf, nil
}
