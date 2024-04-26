package main

import (
	"log"
	"petstore/config"
	"petstore/run"

	"github.com/joho/godotenv"
)

//	@title			HugoMap
//	@version		1.0
//	@description	API Server for HugoMap Application
//
//	@host			localhost:8080
//	@BasePath		/

// @securityDefinitions.apikey	ApiKeyAuth
//
// @in							Login
// @name						Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env files found")
	}
	conf := config.NewAppConf()

	App := run.NewApp(conf)

	err := App.
		// Инициализация
		Bootstrap().
		// Запуск
		Run()

	log.Println(err)
}
