package service

import (
	"context"
	"encoding/json"
	"errors"
	"petstore/internal/infrastructure/db/dao"
	"petstore/internal/models"
	"petstore/internal/modules/pet/storage"

	"github.com/go-chi/jwtauth"
)

type Peter interface {
	CreatePet(pet models.Pet) error
	UpdatePet(pet models.Pet) error
	FindByStatusPet(status string) ([]models.Pet, error)
	FindByIdPet(id string) (models.Pet, error)
	UpdateByIdPet(id, name, status string) error
	DeleteById(api_key, id string) error
}

type Pet struct {
	token *jwtauth.JWTAuth
	storage.PeterRepository
}

func NewPetService(token *jwtauth.JWTAuth, PetRep storage.PeterRepository) Peter {
	return &Pet{
		token:           token,
		PeterRepository: PetRep,
	}
}

func fromPetDbToPet(pet models.PetDB) models.Pet {
	var photo []string
	var tags []models.Tag
	var category models.Category
	json.Unmarshal([]byte(pet.PhotoUrls), &photo)
	json.Unmarshal([]byte(pet.Tags), &tags)
	json.Unmarshal([]byte(pet.Category), &category)

	return models.Pet{
		Id:        pet.Id,
		Category:  category,
		Name:      pet.Name,
		PhotoUrls: photo,
		Tags:      tags,
		Status:    pet.Status,
	}
}

func fromPetToPetDb(pet models.Pet) models.PetDB {
	photo, _ := json.Marshal(pet.PhotoUrls)
	tags, _ := json.Marshal(pet.Tags)
	category, _ := json.Marshal(pet.Category)

	return models.PetDB{
		Id:        pet.Id,
		Category:  string(category),
		Name:      pet.Name,
		PhotoUrls: string(photo),
		Tags:      string(tags),
		Status:    pet.Status,
	}
}

func (p *Pet) CreatePet(pet models.Pet) error {
	return p.PeterRepository.Create(context.Background(), fromPetToPetDb(pet))
}
func (p *Pet) UpdatePet(pet models.Pet) error {
	return p.Update(context.Background(), fromPetToPetDb(pet))
}
func (p *Pet) FindByStatusPet(status string) ([]models.Pet, error) {
	petsDb, err := p.PeterRepository.List(context.Background(), dao.Condition{
		LimitOffset: &dao.LimitOffset{
			Limit:  10000,
			Offset: 0,
		},
		Equal: map[string]interface{}{"status": status},
	})
	if err != nil {
		return nil, err
	}
	pets := make([]models.Pet, 0, len(petsDb))

	for _, pet := range petsDb {
		pets = append(pets, fromPetDbToPet(pet))
	}
	return pets, nil
}
func (p *Pet) FindByIdPet(id string) (models.Pet, error) {
	petsDb, err := p.PeterRepository.GetByID(context.Background(), id)
	if err != nil {
		return models.Pet{}, err
	}
	return fromPetDbToPet(petsDb), nil
}
func (p *Pet) UpdateByIdPet(id, name, status string) error {
	petsDb, err := p.PeterRepository.GetByID(context.Background(), id)
	if err != nil {
		return err
	}
	petsDb.Name = name
	petsDb.Status = status
	return p.PeterRepository.Update(context.Background(), petsDb)
}
func (p *Pet) DeleteById(api_key, id string) error {
	if p.PeterRepository.GetApiKey(context.Background(), api_key) {
		return p.PeterRepository.Delete(context.Background(), id)
	}
	return errors.New("not such api_key")
}
