package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"petstore/internal/infrastructure/responder"
	"petstore/internal/models"
	service "petstore/internal/modules/pet/service"

	"github.com/go-chi/chi"
)

const (
	//status code
	available = "available"
	pending   = "pending"
	sold      = "sold"
)

type Peter interface {
	CreatePetHandler(w http.ResponseWriter, r *http.Request)
	PutPetHandler(w http.ResponseWriter, r *http.Request)
	FindByStatusPetHandler(w http.ResponseWriter, r *http.Request)
	FindByIdPetHandler(w http.ResponseWriter, r *http.Request)
	UpdateByIdPetHandler(w http.ResponseWriter, r *http.Request)
	DeleteByIdHandler(w http.ResponseWriter, r *http.Request)
}

type Pet struct {
	pet     service.Peter
	respond responder.Responder
}

func NewPetHandler(pet service.Peter, respond responder.Responder) Peter {
	return &Pet{
		pet:     pet,
		respond: respond,
	}
}

// @Summary	Add a new pet to the store
// @Tags		pet
// @Security	ApiKeyAuth
// @Accept		json
// @Produce	json
// @Param		request	body	models.Pet	true	"Pet object that needs to be added to the store"
// @Success	200
// @Router		/pet [post]
func (p *Pet) CreatePetHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Pet

	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		p.respond.ErrorBadRequest(w, err)
		return
	}
	err = p.pet.CreatePet(reqBody)
	if err != nil {
		p.respond.ErrorInternal(w, err)
		return
	}
	p.respond.OutputJSON(w, "")
}

// @Summary	Update an existing pet to the store
// @Tags		pet
// @Security	ApiKeyAuth
// @Accept		json
// @Produce	json
// @Param		request	body	models.Pet	true	"Pet object that needs to be added to the store"
// @Success	200
// @Router		/pet [put]
func (p *Pet) PutPetHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Pet

	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.respond.ErrorInternal(w, err)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		p.respond.ErrorBadRequest(w, err)
		return
	}
	err = p.pet.UpdatePet(reqBody)
	if err != nil {
		p.respond.ErrorInternal(w, err)
		return
	}
	p.respond.OutputJSON(w, "")
}

// @Summary	Find Pets by status
// @Tags		pet
// @Security	ApiKeyAuth
// @Accept		json
// @Produce	json
// @Param		status	query		string	true	"Status values that need to be considered for filter\nAvailable values : available, pending, sold"
// @Success	200		{object}	[]models.Pet
// @Router		/pet/findByStatus [get]
func (p *Pet) FindByStatusPetHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == available || status == pending || status == sold {
		pets, err := p.pet.FindByStatusPet(status)
		if err != nil {
			p.respond.ErrorInternal(w, err)
			return
		}
		p.respond.OutputJSON(w, pets)

	} else {
		p.respond.ErrorBadRequest(w, errors.New("not such status"))
		return
	}
}

// @Summary	Find pet by ID
// @Tags		pet
// @Security	ApiKeyAuth
// @Accept		json
// @Produce	json
// @Param		petId	path		int	true	"ID of pet to return"
// @Success	200		{object}	models.Pet
// @Router		/pet/{petId} [get]
func (p *Pet) FindByIdPetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "petId")
	pet, err := p.pet.FindByIdPet(id)
	if err != nil {
		p.respond.ErrorInternal(w, err)
		return
	}
	p.respond.OutputJSON(w, pet)
}

// @Summary	Updates a pet in the store with form data
// @Tags		pet
// @Security	ApiKeyAuth
// @Accept		json
// @Produce	json
// @Param		petId	path	int		true	"ID of pet that needs to be updated"
// @Param		name	query	string	true	"Updated name of the pet"
// @Param		status	query	string	true	"Updated status of the pet"
// @Success	200
// @Router		/pet/{petId} [post]
func (p *Pet) UpdateByIdPetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "petId")
	name := r.URL.Query().Get("name")
	status := r.URL.Query().Get("status")
	if status == available || status == pending || status == sold {
		err := p.pet.UpdateByIdPet(id, name, status)
		if err != nil {
			p.respond.ErrorInternal(w, err)
			return
		}
		p.respond.OutputJSON(w, "")
	} else {
		p.respond.ErrorBadRequest(w, errors.New("not such status"))
		return
	}
}

// @Summary	Dletes a pet
// @Tags		pet
// @Security	ApiKeyAuth
// @Accept		json
// @Produce	json
// @Param		petId	path	int		true	"Pet id to delete"
// @Param		api_key	query	string	true	"api_key"
// @Success	200
// @Router		/pet/{petId} [delete]
func (p *Pet) DeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "petId")
	api_key := r.URL.Query().Get("api_key")
	err := p.pet.DeleteById(api_key, id)
	if err != nil {
		p.respond.ErrorUnauthorized(w, err)
		return
	}
	p.respond.OutputJSON(w, "")
}
