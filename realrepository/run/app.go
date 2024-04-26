package run

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"repository/config"
	"repository/internal/infrastructure/db/dao"
	"repository/internal/infrastructure/db/migrator"
	"repository/internal/models"
	"repository/internal/router"
	"repository/internal/storage"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Application - интерфейс приложения
type Application interface {
	Runner
	Bootstraper
}

// Runner - интерфейс запуска приложения
type Runner interface {
	Run() error
}

// Bootstraper - интерфейс инициализации приложения
type Bootstraper interface {
	Bootstrap(options ...interface{}) Runner
}

type App struct {
	conf    *config.AppConf
	Sig     chan os.Signal
	Storage storage.UserRepository
}

func NewApp(conf *config.AppConf) *App {
	return &App{conf: conf, Sig: make(chan os.Signal, 1)}
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:         ":" + a.conf.Server.Port,
		Handler:      router.NewApiHandler(a.Storage), // Здесь должен быть ваш обработчик запросов
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Ожидание сигнала остановки
	<-stopChan

	// Создание контекста с таймаутом в пять секунд для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Остановка сервера с использованием graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")

	return nil
}

func (a *App) Bootstrap(options ...interface{}) Runner {

	dbRaw, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", a.conf.DB.Host, a.conf.DB.Port, a.conf.DB.User, a.conf.DB.Password, a.conf.DB.Name))
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(dbRaw, "postgres")
	sqlAdapter := dao.NewDAO(dbx)

	var generator migrator.SQLiteGenerator
	m := migrator.NewMigrator(dbRaw, &generator)
	err = m.Migrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	a.Storage = storage.NewUserStorage(sqlAdapter)

	return a
}
