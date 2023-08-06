package usecases

import (
	"errors"
	"log"

	"github.com/olumidayy/go-shawty/src/domain/models"
)


type AuthUsecases struct {
	Repository UserRepository
}


func NewAuthUsecases(r UserRepository) *AuthUsecases {
	return &AuthUsecases{
		Repository: r,
	}
}

func (au* AuthUsecases) RegisterUser (user models.User) (*models.User, error) {
	exists, err := au.Repository.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	log.Println(exists)
	if exists != nil {
		return nil, errors.New("A user with this email already exists.")
	}
	return au.Repository.Create(user)
}
