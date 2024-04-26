package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/ptflp/geotask/run"
)

// @title			HugoMap
// @version		1.0
// @description	API Server for HugoMap Application
//
// @host			localhost:8080
// @BasePath		/
func main() {
	godotenv.Load()
	// инициализация приложения
	app := run.NewApp()
	// запуск приложения
	err := app.Run()
	// в случае ошибки выводим ее в лог и завершаем работу с кодом 2
	if err != nil {
		log.Println(fmt.Sprintf("error: %s", err))
		os.Exit(2)
	}
}
