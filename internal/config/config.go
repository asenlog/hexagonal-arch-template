package config

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

var errConfigPath = errors.New("failed to determine config path")

// HTTP configuration
type HTTP struct {
	Port            int           `env:"HTTP_SERVER_PORT"`
	ShutdownTimeout time.Duration `env:"HTTP_SERVER_SHUTDOWN_TIMEOUT"`
	ReadTimeout     time.Duration `env:"HTTP_SERVER_READ_TIMEOUT"`
	WriteTimeout    time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT"`
}

// DB configuration
type DB struct {
	Username   string `env:"DB_USERNAME"`
	Password   string `env:"DB_PASSWORD"`
	WriterHost string `env:"DB_WRITER_HOST"`
	ReaderHost string `env:"DB_READER_HOST"`
	Port       string `env:"DB_PORT"`
	DB         string `env:"DB_DB"`
	SSLMode    string `env:"DB_SSL_MODE"`
}

// Config holds the service configuration
type Config struct {
	DB
	HTTP
}

// Load loads all the configuration of the app
func Load() (*Config, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("runtime.Caller failed error: %w", errConfigPath)
	}

	configPath := filepath.Dir(file)

	viper.AddConfigPath(fmt.Sprintf("%s/../../", configPath))
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			log.Print("failed to find config file, looking for env vars...")
		} else {
			return nil, fmt.Errorf("failed to parse config: %w", err)
		}
	}

	return &Config{
		// configure the http service
		HTTP: HTTP{
			Port:            viper.GetInt("HTTP_SERVER_PORT"),
			ShutdownTimeout: viper.GetDuration("HTTP_SERVER_SHUTDOWN_TIMEOUT"),
			ReadTimeout:     viper.GetDuration("HTTP_SERVER_READ_TIMEOUT"),
			WriteTimeout:    viper.GetDuration("HTTP_SERVER_WRITE_TIMEOUT"),
		},
		DB: DB{
			Username:   viper.Get("DB_USERNAME").(string),
			Password:   viper.Get("DB_PASSWORD").(string),
			WriterHost: viper.Get("DB_WRITER_HOST").(string),
			ReaderHost: viper.Get("DB_READER_HOST").(string),
			Port:       viper.Get("DB_PORT").(string),
			DB:         viper.Get("DB_DB").(string),
		},
	}, nil
}
