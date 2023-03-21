package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	defaultConfigPath = "../../configs/.env"
	defaultConfigType = "env"
)

type Config struct {
	Reddis  RedisConfig
	Gin     GinConfig
	Http    HTTPConfig
	Limiter LimiterConfig
}

type GinConfig struct {
	GinMode string
}

type RedisConfig struct {
	Addr     string
	Password string
}

type HTTPConfig struct {
	Host               string
	Port               string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

type LimiterConfig struct {
	RPS   int
	Burst int
	TTL   time.Duration
}

func NewConfig() (*Config, error) {
	viper.SetConfigType(defaultConfigType)
	viper.SetConfigFile(defaultConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	viper.AutomaticEnv()

	return &Config{
		Gin: GinConfig{
			GinMode: viper.GetString("GIN_MODE"),
		},
		Http: HTTPConfig{
			Host:               "",
			Port:               viper.GetString("HTTP_PORT"),
			ReadTimeout:        time.Duration(viper.GetInt("HTTP_READ_TIMEOUT")) * time.Second,
			WriteTimeout:       time.Duration(viper.GetInt("HTTP_WRITE_TIMEOUT")) * time.Second,
			MaxHeaderMegabytes: viper.GetInt("HTTP_MAX_HEADER_MEGABYTES"),
		},
		Limiter: LimiterConfig{
			RPS:   viper.GetInt("LIMITER_RPS"),
			Burst: viper.GetInt("LIMITER_BURST"),
			TTL:   time.Duration(viper.GetInt("LIMITER_TTL")) * time.Second,
		},
		Reddis: RedisConfig{
			Addr:     viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDDIS_PASSWORD"),
		},
	}, nil
}
