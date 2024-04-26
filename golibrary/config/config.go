package config

import (
	"os"
	"time"
)

const (
	shutDownTime = 5 * time.Second
)

type AppConf struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

type DB struct {
	Name     string `yaml:"name"`
	User     string `json:"-" yaml:"user"`
	Password string `json:"-" yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type Server struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

func NewAppConf() AppConf {
	return AppConf{
		Server: Server{
			Port:            os.Getenv("SERVER_PORT"),
			ShutdownTimeout: shutDownTime,
		},
		DB: DB{
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}
