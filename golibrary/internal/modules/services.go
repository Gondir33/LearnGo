package modules

import (
	"golibrary/internal/infrastructure/component"
	lService "golibrary/internal/modules/library/service"
	uService "golibrary/internal/modules/user/service"
	"golibrary/internal/storages"
)

type Services struct {
	uService.Userer
	lService.LibraryServicere
}

func NewServices(storages *storages.Storages, components *component.Components) *Services {
	return &Services{
		Userer:           uService.NewUserService(storages.UsererRepository),
		LibraryServicere: lService.NewUserService(storages.LibraryRepository),
	}
}
