package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"petstore/internal/models"
	storage "petstore/internal/modules/user/storage"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth"
)

type Userer interface {
	Create(user models.User) error
	Login(username string, password string) (string, map[string]interface{}, string, error)
	Update(username string, user models.User) error
	Delete(username string) error
	CreateWithArray(users []models.User) error
	GetUser(username string) (models.User, error)
	CheckTokenLogout(token string) bool
	Logout(string)
}

type User struct {
	tokensLogout map[string]bool
	token        *jwtauth.JWTAuth
	storage.UsererRepository
}

func NewUserService(token *jwtauth.JWTAuth, UsererRep storage.UsererRepository) *User {
	return &User{
		tokensLogout:     make(map[string]bool, 10),
		token:            token,
		UsererRepository: UsererRep,
	}
}

func fromUserToUserDB(user models.User) models.UserDB {
	return models.UserDB{
		Id:         user.Id,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Password:   user.Password,
		Phone:      user.Phone,
		UserStatus: user.UserStatus,
	}
}

func fromUserDBToUser(user models.UserDB) models.User {
	return models.User{
		Id:         user.Id,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Password:   user.Password,
		Phone:      user.Phone,
		UserStatus: user.UserStatus,
	}
}

func (u *User) Create(user models.User) error {
	userDb := fromUserToUserDB(user)

	return u.UsererRepository.Create(context.Background(), userDb)
}

func (u *User) Login(username string, password string) (string, map[string]interface{}, string, error) {
	userDb, err := u.GetUser(username)
	if err != nil {
		return "", nil, "", err
	}
	if userDb.Password == password {
		claims := map[string]interface{}{
			"username":   username,
			"exp":        time.Now().Add(time.Hour * 4).UnixNano(),
			"rate_limit": 10,
		}
		_, tokenString, _ := u.token.Encode(claims)

		api_key, err := generateAPIKey(30)

		if err != nil {
			return "", nil, "", err
		}
		err = u.UsererRepository.PutApiKey(context.Background(), models.Api_keyDB{Api_key: api_key})
		if err != nil {
			return "", nil, "", err
		}
		return tokenString, claims, api_key, nil
	}
	return "", nil, "", err
}

func (u *User) Update(username string, user models.User) error {

	userDb, err := u.UsererRepository.GetByUsername(context.Background(), username)
	if err != nil {
		return err
	}
	userDb.Id = user.Id
	userDb.Username = user.Username
	userDb.FirstName = user.FirstName
	userDb.LastName = user.LastName
	userDb.Email = user.Email
	userDb.Password = user.Password
	userDb.Phone = user.Phone
	userDb.UserStatus = user.UserStatus
	return u.UsererRepository.Update(context.Background(), userDb)
}
func (u *User) Delete(username string) error {
	userDb, err := u.UsererRepository.GetByUsername(context.Background(), username)
	if err != nil {
		return err
	}

	return u.UsererRepository.Delete(context.Background(), strconv.Itoa(userDb.Id))
}
func (u *User) CreateWithArray(users []models.User) error {
	for _, user := range users {
		err := u.Create(user)
		if err != nil {
			return err
		}
	}
	return nil
}
func (u *User) GetUser(username string) (models.User, error) {
	userDb, err := u.UsererRepository.GetByUsername(context.Background(), username)
	if err != nil {
		return models.User{}, err
	}
	return fromUserDBToUser(userDb), nil
}

func (u *User) CheckTokenLogout(token string) bool {
	return u.tokensLogout[token]
}

func (u *User) Logout(tokenString string) {
	u.tokensLogout[tokenString] = true
}

func generateAPIKey(length int) (string, error) {
	byteSize := length * 3 / 4
	if length%4 != 0 {
		byteSize++
	}

	keyBytes := make([]byte, byteSize)

	// Read random bytes into the byte slice
	if _, err := rand.Read(keyBytes); err != nil {
		return "", err
	}

	apiKey := base64.URLEncoding.EncodeToString(keyBytes)

	apiKey = apiKey[:length]

	return apiKey, nil
}
