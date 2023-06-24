package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *Config

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DatabaseName string `mapstructure:"databasename"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type Config struct {
	Database struct {
		Master DatabaseConfig   `mapstructure:"master"`
		Slaves []DatabaseConfig `mapstructure:"slaves"`
	}
	ServiceName string `mapstructure:"servicename"`
	Server ServerConfig `mapstructure:"server"` 
	Redis []RedisConfig `mapstructure:"redis"`
}

func GetConfig() *Config {
	return cfg
}

func LoadConfig() error {
	// Set the path to the configuration file
	viper.SetConfigFile("config/test/config.yml")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file: %s", err)
	}

	// Unmarshal the configuration values into the Config struct
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file: %s", err)
	}
	return nil
}
