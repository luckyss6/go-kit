package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Debug    bool     `json:"debug"`
	Redis    Redis    `yaml:"redis"`
	Postgres Postgres `yaml:"postgres"`
	Mysql    Mysql    `yaml:"mysql"`
	Binance  Binance  `yaml:"binance"`
}

type Binance struct {
	ApiKey    string `yaml:"api_key"`
	SecretKey string `yaml:"secret_key"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBname   string `yaml:"dbname"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var cfg = &Config{}

func New() *Config {
	cfgPath := os.Getenv("ENV_PATH")
	if strings.Compare(cfgPath, "") == 0 {
		cfgPath = "."
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfgPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config err: %s", err.Error()))
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		panic(fmt.Errorf("unmarshal cfg err: %s", err.Error()))
	}
	return cfg
}
