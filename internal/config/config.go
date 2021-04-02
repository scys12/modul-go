package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver    string `mapstructure:"db_driver"`
	DBName      string `mapstructure:"db_name"`
	DBUser      string `mapstructure:"db_user"`
	DBPassword  string `mapstructure:"db_password"`
	DBHost      string `mapstructure:"db_host"`
	DBPort      string `mapstructure:"db_port"`
	MaxLifetime int    `mapstructure:"max_lifetime"`
	IdleConn    int    `mapstructure:"idle_conn"`
	OpenConn    int    `mapstructure:"open_conn"`
	JWTSecret   string `mapstructure:"jwt_secret"`
	JWTExpires  int    `mapstructure:"jwt_expires"`
}

func InitConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	if err = viper.Unmarshal(&config); err != nil {
		return
	}
	return
}
