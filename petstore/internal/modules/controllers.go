package modules

import (
	"petstore/internal/infrastructure/db/dao"
	"petstore/internal/infrastructure/responder"
	pHandler "petstore/internal/modules/pet/handler"
	pService "petstore/internal/modules/pet/service"
	pStorage "petstore/internal/modules/pet/storage"
	sHandler "petstore/internal/modules/store/handler"
	sService "petstore/internal/modules/store/service"
	sStorage "petstore/internal/modules/store/storage"
	uHandler "petstore/internal/modules/user/handler"
	uService "petstore/internal/modules/user/service"
	uStorage "petstore/internal/modules/user/storage"

	"github.com/go-chi/jwtauth"
)

type Handlers struct {
	Uhandler uHandler.Userer
	SHandler sHandler.Storere
	PHandler pHandler.Peter
}

func NewHandlers(token *jwtauth.JWTAuth, db dao.IfaceDAO, respond responder.Responder) *Handlers {
	uhandler := uHandler.NewUserHandler(uService.NewUserService(token, uStorage.NewUserStorage(db)), respond)
	sHandler := sHandler.NewStoreHandler(sService.NewStoreService(token, sStorage.NewStoreStorage(db)), respond)
	pHandler := pHandler.NewPetHandler(pService.NewPetService(token, pStorage.NewPetStorage(db)), respond)
	return &Handlers{
		Uhandler: uhandler,
		SHandler: sHandler,
		PHandler: pHandler,
	}
}
