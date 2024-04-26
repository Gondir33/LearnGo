package storages

import (
	lStorage "golibrary/internal/modules/library/storage"
	uStorage "golibrary/internal/modules/user/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storages struct {
	uStorage.UsererRepository
	lStorage.LibraryRepository
}

func NewStorages(pool *pgxpool.Pool) *Storages {
	return &Storages{
		UsererRepository:  uStorage.NewUserStorage(pool),
		LibraryRepository: lStorage.NewLibraryStorage(pool),
	}
}
