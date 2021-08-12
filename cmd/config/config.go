package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DbDrivername string
	Username     string
	Password     string
	Dbname       string
	Sslmode      string
	Network      string
	Address      string
}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}

func GetConfig() *Config {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	return &Config{
		DbDrivername: viper.GetString("db.drivername"),
		Username:     viper.GetString("db.username"),
		Password:     viper.GetString("db.password"),
		Dbname:       viper.GetString("db.dbname"),
		Sslmode:      viper.GetString("db.sslmode"),
		Network:      viper.GetString("server.network"),
		Address:      viper.GetString("server.address"),
	}
}

func (conf Config) GetDbConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		conf.Username, conf.Password, conf.Dbname, conf.Sslmode)
}
