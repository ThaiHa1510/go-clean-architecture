package config

import (
	"fmt"
	"github.com/spf13/viper"
)
type Config struct {
	App: App
	Db: Db
}

type App struct {
	Name: string
	Host: string
	Port: int
}

type Db struct {
	Host: string
	Port: int
	User: string
	Password: string
	DBName: string
	SSLMode: bool
	TimeZone: string
}

func GetConfig() *Config{
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadConfig()
	if err != nil{
		panic("Error when read config")
	}
	return &Config{
		App:{
			Name: viper.GetString("app.name"),
			Host: viper.GetString("app.host"),
			Port: viper.GetInt("app.port")
		},
		Database:{
			Host: viper.GetString("database.host"),
			Port: viper.GetInt("database.port"),
			UserName: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			DbName: viper.GetString("database.dbname"),
			SSLMode: viper.GetBool("database.ssl_mode"),
			TimeZone: viper.GetString("database.time_zone")
		}
	}
}