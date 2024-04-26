package storage

import (
	"context"
	"golibrary/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsererRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, id int) (models.UserDTO, error)
	List(ctx context.Context) ([]models.UserDTO, error)
	GetByUsername(ctx context.Context, username string) (models.UserDTO, error)
}
type UserStorage struct {
	pool *pgxpool.Pool
}

func NewUserStorage(pool *pgxpool.Pool) UsererRepository {
	return &UserStorage{pool}
}

func (s *UserStorage) Create(ctx context.Context, user models.User) error {
	sql := "INSERT INTO users (username) VALUES ($1)"
	rows, err := s.pool.Query(ctx, sql, user.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (s *UserStorage) GetByUsername(ctx context.Context, username string) (models.UserDTO, error) {
	var UserDTO models.UserDTO

	sql := "SELECT id, username FROM users WHERE username=$1"
	if err := s.pool.QueryRow(ctx, sql, username).Scan(&UserDTO.Id, &UserDTO.Name); err != nil {
		return models.UserDTO{}, err
	}
	return UserDTO, nil
}

func (s *UserStorage) GetByID(ctx context.Context, id int) (models.UserDTO, error) {
	var UserDTO models.UserDTO

	sql := "SELECT id, username FROM users WHERE id=$1"
	if err := s.pool.QueryRow(ctx, sql, id).Scan(&UserDTO.Id, &UserDTO.Name); err != nil {
		return models.UserDTO{}, err
	}
	return UserDTO, nil
}

func (s *UserStorage) List(ctx context.Context) ([]models.UserDTO, error) {
	var users []models.UserDTO

	sql := "SELECT id, username FROM users"
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return []models.UserDTO{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserDTO

		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
