package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ListenAddr  string `mapstructure:"LISTEN_ADDR"`
	ListenPort  string `mapstructure:"LISTEN_PORT"`
	Env         string `mapstructure:"ENV"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	RedisUrl    string `mapstructure:"REDIS_URL"`
	RedisPass   string `mapstructure:"REDIS_PASS"`
}

func (c Config) ListenAddrAndPort() string {
	return fmt.Sprintf("%s:%s", c.ListenAddr, c.ListenPort)
}

func FromEnv() (*Config, error) {
	v := viper.New()
	v.SetDefault("LISTEN_ADDR", "0.0.0.0")
	v.SetDefault("LISTEN_PORT", "8080")
	v.SetDefault("ENV", "local")
	v.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/magbat_game?sslmode=disable")
	v.SetDefault("REDIS_URL", "localhost:6379")
	v.SetDefault("REDIS_PASS", "45554555")
	v.SetConfigName("env")
	v.SetConfigFile("../.env")
	_ = v.ReadInConfig()
	v.AutomaticEnv()

	cfg := Config{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
