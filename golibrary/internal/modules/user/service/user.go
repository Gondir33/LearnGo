package service

import (
	"context"
	"golibrary/internal/models"
	"golibrary/internal/modules/user/storage"
)

type Userer interface {
	Create(ctx context.Context, user models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.UserDTO, error)
	GetUserById(ctx context.Context, id int) (models.UserDTO, error)
	List(ctx context.Context) ([]models.UserDTO, error)
}

type User struct {
	storage.UsererRepository
}

func NewUserService(UsererRep storage.UsererRepository) Userer {
	return &User{
		UsererRepository: UsererRep,
	}
}

func (u *User) Create(ctx context.Context, user models.User) error {
	return u.UsererRepository.Create(context.Background(), user)
}

func (u *User) GetUserByUsername(ctx context.Context, username string) (models.UserDTO, error) {
	return u.UsererRepository.GetByUsername(ctx, username)
}

func (u *User) List(ctx context.Context) ([]models.UserDTO, error) {
	return u.UsererRepository.List(ctx)
}
func (u *User) GetUserById(ctx context.Context, id int) (models.UserDTO, error) {
	return u.UsererRepository.GetByID(ctx, id)
}
