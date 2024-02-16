package configs

import (
	"log"

	"github.com/spf13/viper"
)


type EnvsStruct struct {
	MongoURI string
	BaseUrl string
	UserServiceApiKey string
	Env string
	Port string
	UserServiceBaseURI string
}

var Envs = &EnvsStruct{}

func LoadEnv() (*EnvsStruct, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
		  return nil, err
    }

	 Envs.MongoURI = viper.GetString("MONGO_URI")
	 Envs.BaseUrl = viper.GetString("BASE_URL")
	 Envs.Env = viper.GetString("ENV")
	 Envs.Port = viper.GetString("PORT")
	 Envs.UserServiceApiKey = viper.GetString("USER_SERVICE_API_KEY")
	 Envs.UserServiceBaseURI = viper.GetString("USER_SERVICE_BASE_URI")

	 return Envs, nil
}