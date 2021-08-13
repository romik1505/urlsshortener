package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	DbDrivername string
	Host         string
	Port         string
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
		Host:         os.Getenv("POSTGRES_HOST"),
		Port:         viper.GetString("db.port"),
		Username:     viper.GetString("db.username"),
		Password:     viper.GetString("db.password"),
		Dbname:       viper.GetString("db.dbname"),
		Sslmode:      viper.GetString("db.sslmode"),
		Network:      viper.GetString("server.network"),
		Address:      viper.GetString("server.address"),
	}
}

func (conf Config) GetDbConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host, conf.Port, conf.Username, conf.Dbname, conf.Password, conf.Sslmode)
}
