package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Db struct{
		Drivername string
		User string
		Password string
		Dbname string
		Sslmode string
	}
	Server struct{
		Network string
		Address string
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func (conf Config) GetDbConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		conf.Db.User, conf.Db.Password, conf.Db.Dbname, conf.Db.Sslmode)
}