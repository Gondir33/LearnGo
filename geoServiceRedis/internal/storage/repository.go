package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"proxy/internal/models"
	"time"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var SomeReposProxy *SomeRepositoryProxy
var PoolPgx *pgxpool.Pool

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no .env files found")
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("postgres"))
	PoolPgx, err = pgxpool.Connect(context.TODO(), dsn)
	if err != nil {
		log.Fatal("connect to db")
	}
	SomeReposProxy = NewSomeRepositoryProxy(PoolPgx)
}

type SomeRepository interface {
	GetData(ctx context.Context, query string) ([]models.Address, error)
}

type SomeRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewSomeRepository() *SomeRepositoryImpl {
	return &SomeRepositoryImpl{pool: PoolPgx}
}

func (r *SomeRepositoryImpl) GetData(ctx context.Context, query string) ([]models.Address, error) {
	return CheckLevenshtein(ctx, query)
}

type SomeRepositoryProxy struct {
	repository SomeRepository
	cache      *redis.Client
}

func NewSomeRepositoryProxy(pool *pgxpool.Pool) *SomeRepositoryProxy {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &SomeRepositoryProxy{
		repository: NewSomeRepository(),
		cache:      client,
	}
}

func (r *SomeRepositoryProxy) GetData(ctx context.Context, query string) ([]models.Address, error) {
	addressesAns := make([]models.Address, 0)
	// Здесь происходит проверка наличия данных в кэше
	cmd := r.cache.Get(query)
	// Если данные есть в кэше, то они возвращаются
	if cmd.Err() == nil {
		str, err := cmd.Result()
		if err != redis.Nil {
			err := json.Unmarshal([]byte(str), &addressesAns)
			return addressesAns, err
		}
	}
	// Если данных нет в кэше, то они запрашиваются у оригинального объекта и сохраняются в кэш

	creds := client.Credentials{
		ApiKeyValue:    "0eade739b4a52041d493615d92913bfa4a2ebcab",
		SecretKeyValue: "2a6a465a6d34c9ad838cf089cb3cc323ab2fa296",
	}
	api := dadata.NewSuggestApi(client.WithCredentialProvider(&creds))

	params := suggest.RequestParams{
		Query: query,
	}

	addresses, err := api.Address(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	for _, address := range addresses {
		addressesAns = append(addressesAns, models.Address{Lat: address.Data.GeoLat, Lng: address.Data.GeoLon})
	}

	str, err := json.Marshal(addressesAns)
	if err != nil {
		return nil, err
	}
	err = r.cache.Set(query, str, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return addressesAns, nil
}

func CheckLevenshtein(ctx context.Context, query string) ([]models.Address, error) {
	var id int
	queryDB := "SELECT id FROM search_history WHERE similarity(search_request, $1) > 0.7 ORDER BY similarity(search_request, $1) DESC"
	if err := PoolPgx.QueryRow(ctx, queryDB, query).Scan(&id); err != nil {
		return []models.Address{}, err
	}

	queryDBId := "SELECT id_address FROM history_search_address WHERE id_search_history = $1"
	rows, err := PoolPgx.Query(ctx, queryDBId, id)
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
		if err = PoolPgx.QueryRow(ctx, queryDbAddres, id).Scan(&lat, &lon); err != nil {
			return adressess, err
		}
		adressess = append(adressess, models.Address{Lat: lat, Lng: lon})
	}
	return adressess, nil
}

func AddCache(ctx context.Context, addresses []models.Address, query string) error {
	var idSearch_history int
	queryDbSerachHistory := "INSERT INTO search_history (search_request) VALUES ($1) RETURNING id"
	if err := PoolPgx.QueryRow(ctx, queryDbSerachHistory, query).Scan(&idSearch_history); err != nil {
		return err
	}
	for _, address := range addresses {
		var idAddress int

		queryDbAddress := "INSERT INTO addresses (lat, lon) VALUES ($1,$2) RETURNING id"
		if err := PoolPgx.QueryRow(ctx, queryDbAddress, address.Lat, address.Lng).Scan(&idAddress); err != nil {
			return err
		}
		queryDbHistorySearchAddress := "INSERT INTO history_search_address (id_search_history, id_address) VALUES ($1,$2)"
		PoolPgx.QueryRow(ctx, queryDbHistorySearchAddress, idSearch_history, idAddress)
	}
	return nil
}
