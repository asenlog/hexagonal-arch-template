package config

import (
	"time"

	"github.com/spf13/viper"
)

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
}

// Config holds the service configuration
type Config struct {
	DB
	HTTP
}

// Load loads all the configuration of the app
func Load() (*Config, error) {
	viper.AutomaticEnv()
	
	return &Config{
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
