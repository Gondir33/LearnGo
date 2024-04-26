package modules

import (
	"golibrary/internal/infrastructure/component"
	lHandler "golibrary/internal/modules/library/controller"
	uHandler "golibrary/internal/modules/user/controller"
)

type Controllers struct {
	uHandler.Userer
	lHandler.Libraryer
}

func NewControllers(services *Services, components *component.Components) *Controllers {
	return &Controllers{
		Userer:    uHandler.NewUserHandler(services.Userer, components.Responder, components.Decoder),
		Libraryer: lHandler.NewLibraryHandler(services.LibraryServicere, components.Responder, components.Decoder),
	}
}
