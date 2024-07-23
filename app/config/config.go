package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	Name string
	Env  string
	Port string
	Host string
}

type DbConfig struct {
	Host        string
	Port        string
	Dbname      string
	Username    string
	Password    string
	DbIsMigrate bool
}
type JwtConfig struct {
	Secret string
}
type Configs struct {
	Appconfig AppConfig
	Dbconfig  DbConfig
	Jwtconfig JwtConfig
}

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	if configs == nil {
		lock.Lock()

		configs = &Configs{
			Appconfig: AppConfig{
				Name: os.Getenv("APP_NAME"),
				Env:  os.Getenv("APP_ENV"),
				Port: os.Getenv("APP_PORT"),
				Host: os.Getenv("APP_HOST"),
			},
			Dbconfig: DbConfig{
				Host:        os.Getenv("DB_HOST"),
				Port:        os.Getenv("DB_PORT"),
				Dbname:      os.Getenv("DB_NAME"),
				Username:    os.Getenv("DB_USER"),
				Password:    os.Getenv("DB_PASS"),
				DbIsMigrate: os.Getenv("DB_ISMIGRATE") == "true",
			},
			Jwtconfig: JwtConfig{
				Secret: os.Getenv("JWT_SECRET"),
			},
		}
		lock.Unlock()
	}

	return configs
}
