package main

import (
	"golibrary/config"
	"golibrary/run"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//	@title			HugoMap
//	@version		1.0
//	@description	API Server for HugoMap Application
//
//	@host			localhost:8080
//	@BasePath		/

//	@in		Login
//	@name	Authorization
func main() {
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env files found")
	}
	// Создаем конфигурацию приложения
	conf := config.NewAppConf()

	// Создаем инстанс приложения
	App := run.NewApp(conf)

	exitCode := App.
		// Инициализация
		Bootstrap().
		// Запуск
		Run()

	os.Exit(exitCode)
}
