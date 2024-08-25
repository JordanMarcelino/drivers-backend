package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App        AppConfig
	HttpServer HttpServerConfig
	Database   DatabaseConfig
}

type AppConfig struct {
	Environment string `mapstructure:"APP_ENVIRONMENT"`
}

type HttpServerConfig struct {
	Host                 string `mapstructure:"HTTP_SERVER_HOST"`
	Port                 int    `mapstructure:"HTTP_SERVER_PORT"`
	GracePeriod          int    `mapstructure:"HTTP_SERVER_GRACE_PERIOD"`
	RequestTimeoutPeriod int    `mapstructure:"HTTP_SERVER_REQUEST_TIMEOUT_PERIOD"`
}

type DatabaseConfig struct {
	Host                  string `mapstructure:"DB_HOST"`
	DbName                string `mapstructure:"DB_NAME"`
	Username              string `mapstructure:"DB_USER"`
	Password              string `mapstructure:"DB_PASSWORD"`
	Sslmode               string `mapstructure:"DB_SSL_MODE"`
	Port                  int    `mapstructure:"DB_PORT"`
	MaxIdleConn           int    `mapstructure:"DB_MAX_IDLE_CONN"`
	MaxOpenConn           int    `mapstructure:"DB_MAX_OPEN_CONN"`
	MaxConnLifetimeMinute int    `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

func InitConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	return &Config{
		App:        initAppConfig(),
		Database:   initDbConfig(),
		HttpServer: initHttpServerConfig(),
	}
}

func initDbConfig() DatabaseConfig {
	dbConfig := DatabaseConfig{}

	if err := viper.Unmarshal(&dbConfig); err != nil {
		log.Fatalf("error mapping database config: %v", err)
	}

	return dbConfig
}

func initHttpServerConfig() HttpServerConfig {
	httpServerConfig := HttpServerConfig{}

	if err := viper.Unmarshal(&httpServerConfig); err != nil {
		log.Fatalf("error mapping http server config: %v", err)
	}

	return httpServerConfig
}

func initAppConfig() AppConfig {
	appConfig := AppConfig{}

	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("error mapping app config: %v", err)
	}

	return appConfig
}
