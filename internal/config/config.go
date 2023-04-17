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
	Auth       *AuthConfig
	Reddis     *RedisConfig
	Mongo      *MongoConfig
	Gin        *GinConfig
	Http       *HTTPConfig
	Limiter    *LimiterConfig
	Blockchain *BlockchainConfig
}

type AuthConfig struct {
	JWT          JWTConfig
	PasswordSalt string
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	SigningKey      string
}

type GinConfig struct {
	GinMode string
}

type MongoConfig struct {
	URI      string
	User     string
	Password string
	Name     string
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

type BlockchainConfig struct {
	AlchemyUrl string
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
		Auth: &AuthConfig{
			PasswordSalt: viper.GetString("PASSWORD_SALT"),
			JWT: JWTConfig{
				SigningKey:      viper.GetString("JWT_SIGNING_KEY"),
				AccessTokenTTL:  time.Duration(viper.GetInt("ACCESS_TOKEN_TTL")) * time.Second,
				RefreshTokenTTL: time.Duration(viper.GetInt("REFRESH_TOKEN_TTL")) * time.Second,
			},
		},
		Gin: &GinConfig{
			GinMode: viper.GetString("GIN_MODE"),
		},
		Http: &HTTPConfig{
			Host:               "",
			Port:               viper.GetString("HTTP_PORT"),
			ReadTimeout:        time.Duration(viper.GetInt("HTTP_READ_TIMEOUT")) * time.Second,
			WriteTimeout:       time.Duration(viper.GetInt("HTTP_WRITE_TIMEOUT")) * time.Second,
			MaxHeaderMegabytes: viper.GetInt("HTTP_MAX_HEADER_MEGABYTES"),
		},
		Limiter: &LimiterConfig{
			RPS:   viper.GetInt("LIMITER_RPS"),
			Burst: viper.GetInt("LIMITER_BURST"),
			TTL:   time.Duration(viper.GetInt("LIMITER_TTL")) * time.Second,
		},
		Reddis: &RedisConfig{
			Addr:     viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDDIS_PASSWORD"),
		},
		Mongo: &MongoConfig{
			URI:      viper.GetString("MONGODB_URI"),
			User:     viper.GetString("MONGODB_USERNAME"),
			Password: viper.GetString("MONGODB_PASSWORD"),
			Name:     viper.GetString("MONGODB_DATABASE"),
		},
		Blockchain: &BlockchainConfig{
			AlchemyUrl: viper.GetString("ALCHEMY_URL"),
		},
	}, nil
}
