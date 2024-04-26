package run

import (
	"context"
	"fmt"
	"golibrary/config"
	"golibrary/internal/errors"
	"golibrary/internal/infrastructure/component"
	"golibrary/internal/infrastructure/responder"
	"golibrary/internal/infrastructure/router"
	"golibrary/internal/infrastructure/server"
	"golibrary/internal/models"
	"golibrary/internal/modules"
	"golibrary/internal/storages"
	"log"
	"net/http"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// Application - интерфейс приложения
type Application interface {
	Runner
	Bootstraper
}

// Runner - интерфейс запуска приложения
type Runner interface {
	Run() int
}

// Bootstraper - интерфейс инициализации приложения
type Bootstraper interface {
	Bootstrap(options ...interface{}) Runner
}

type App struct {
	conf   config.AppConf
	logger *zap.Logger

	srv      server.Server
	Sig      chan os.Signal
	Storages *storages.Storages
	Servises *modules.Services
}

func NewApp(conf config.AppConf) *App {
	return &App{conf: conf, Sig: make(chan os.Signal, 1)}
}

func (a *App) Run() int {
	// на русском
	// создаем контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	errGroup, ctx := errgroup.WithContext(ctx)

	// запускаем горутину для graceful shutdown
	// при получении сигнала SIGINT
	// вызываем cancel для контекста
	errGroup.Go(func() error {
		sigInt := <-a.Sig
		a.logger.Info("signal interrupt recieved", zap.Stringer("os_signal", sigInt))
		cancel()
		return nil
	})

	errGroup.Go(func() error {
		err := a.srv.Serve(ctx)
		if err != nil && err != http.ErrServerClosed {
			a.logger.Error("app: server error", zap.Error(err))
			return err
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return errors.GeneralError
	}
	return errors.NoError
}

func (a *App) Bootstrap(options ...interface{}) Runner {
	// инициализация логгера
	logger, _ := zap.NewProduction()
	a.logger = logger
	// инициализация декодера
	decoder := godecoder.NewDecoder()
	// инициализация менеджера ответов сервера
	responseManager := responder.NewResponder(decoder, logger)
	// инициализация компонентов
	components := component.NewComponents(a.conf, responseManager, decoder, a.logger)

	// инициализация базы данных sql и его адаптера
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		a.conf.DB.User, a.conf.DB.Password, a.conf.DB.Host, a.conf.DB.Port, a.conf.DB.Name)
	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		a.logger.Fatal("error init db", zap.Error(err))
	}

	// инициализация хранилищ
	newStorages := storages.NewStorages(pool)

	a.Storages = newStorages
	//генерация данных
	GenerationData(newStorages, pool)
	// инициализация сервисов
	services := modules.NewServices(newStorages, components)
	a.Servises = services
	controllers := modules.NewControllers(services, components)
	// инициализация роутера
	r := router.NewRouter(controllers, components)
	// конфигурация сервера
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler: r,
	}
	a.srv = server.NewHttpServer(a.conf.Server, srv, a.logger)
	return a
}

func GenerationData(storages *storages.Storages, pool *pgxpool.Pool) {
	var count int

	sql := "SELECT COUNT(*) FROM author"
	err := pool.QueryRow(context.Background(), sql).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		if err := GenerateAuthors(storages); err != nil {
			log.Fatal(err)
		}
	}
	count = 0
	sql = "SELECT COUNT(*) FROM book"
	err = pool.QueryRow(context.Background(), sql).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		if err := GenerateBooks(storages); err != nil {
			log.Fatal(err)
		}
	}

	count = 0
	sql = "SELECT COUNT(*) FROM users"
	err = pool.QueryRow(context.Background(), sql).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		if err := GenerateUsers(storages); err != nil {
			log.Fatal(err)
		}
	}
}

func GenerateAuthors(storages *storages.Storages) error {
	for i := 0; i < 10; i++ {
		name := gofakeit.Book().Author
		_, err := storages.LibraryRepository.CreateAuthor(context.Background(), models.Author{Name: name})
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateBooks(storages *storages.Storages) error {
	for i := 0; i < 100; i++ {
		nameBook := gofakeit.Book().Title
		id_book, err := storages.LibraryRepository.CreateBook(context.Background(), models.Book{Name: nameBook})
		if err != nil {
			return err
		}
		id_author := gofakeit.IntRange(1, 10)
		err = storages.LibraryRepository.CreateAuthorBook(context.Background(), id_author, id_book)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateUsers(storages *storages.Storages) error {
	for i := 0; i < 51; i++ {
		name := gofakeit.Name()
		err := storages.UsererRepository.Create(context.Background(), models.User{Name: name})
		if err != nil {
			return err
		}
	}
	return nil
}
