package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql `toml:"mysql"`
	Redis `toml:"redis"`
}
type Mysql struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DB       string `toml:"db"`
}

type Redis struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

func NewCfg() *Config {
	cfgPath := os.Getenv("ENV_PATH")
	if strings.Compare(cfgPath, "") == 0 {
		cfgPath = "./"
	}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(cfgPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config err: %s", err.Error()))
	}

	cfg := new(Config)

	err = v.Unmarshal(cfg)
	if err != nil {
		panic(fmt.Errorf("unmarshal cfg err: %s", err.Error()))
	}
	return cfg
}
