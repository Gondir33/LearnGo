package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"proxy/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no .env files found")
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("postgres"))
	pool, err = pgxpool.Connect(context.TODO(), dsn)
	if err != nil {
		log.Fatal("connect to db")
	}
}

func CheckLevenshtein(ctx context.Context, query string) ([]models.Address, error) {
	var id int
	queryDB := "SELECT id FROM search_history WHERE similarity(search_request, $1) > 0.7 ORDER BY similarity(search_request, $1) DESC"
	if err := pool.QueryRow(ctx, queryDB, query).Scan(&id); err != nil {
		return []models.Address{}, err
	}

	queryDBId := "SELECT id_address FROM history_search_address WHERE id_search_history = $1"
	rows, err := pool.Query(ctx, queryDBId, id)
	if err != nil {
		return []models.Address{}, err
	}
	defer rows.Close()

	adressess := make([]models.Address, 0)
	for rows.Next() {
		var lat string
		var lon string
		if err = rows.Scan(&id); err != nil {
			return adressess, err
		}
		queryDbAddres := "SELECT lat, lon FROM addresses WHERE id = $1"
		if err = pool.QueryRow(ctx, queryDbAddres, id).Scan(&lat, &lon); err != nil {
			return adressess, err
		}
		adressess = append(adressess, models.Address{Lat: lat, Lng: lon})
	}
	return adressess, nil
}

func AddCache(ctx context.Context, addresses []models.Address, query string) error {
	var idSearch_history int
	queryDbSerachHistory := "INSERT INTO search_history (search_request) VALUES ($1) RETURNING id"
	if err := pool.QueryRow(ctx, queryDbSerachHistory, query).Scan(&idSearch_history); err != nil {
		return err
	}
	for _, address := range addresses {
		var idAddress int

		queryDbAddress := "INSERT INTO addresses (lat, lon) VALUES ($1,$2) RETURNING id"
		if err := pool.QueryRow(ctx, queryDbAddress, address.Lat, address.Lng).Scan(&idAddress); err != nil {
			return err
		}
		queryDbHistorySearchAddress := "INSERT INTO history_search_address (id_search_history, id_address) VALUES ($1,$2)"
		pool.QueryRow(ctx, queryDbHistorySearchAddress, idSearch_history, idAddress)
	}
	return nil
}
