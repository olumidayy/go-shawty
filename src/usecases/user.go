package usecases

import (
	"github.com/olumidayy/go-shawty/src/domain/models"
	"github.com/olumidayy/go-shawty/src/domain/repositories"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) uint
	GetAll() []models.User
	GetById(id uint) models.User
	GetOne(data interface{}) models.User
	Update(id uint, data interface{}) models.User
	Delete(id uint)
}

type UserUsecases struct {
	Repository UserRepository
}

func NewUserUsecases(DB *gorm.DB) *UserUsecases {
	return &UserUsecases{
		Repository: repositories.NewUserRepository(DB),
	}
}

func (uu* UserUsecases) CreateNewUser (user models.User) uint {
	userID := uu.Repository.Create(user)
	return userID
}

func (uu* UserUsecases) GetAllUsers () []models.User {
	var users []models.User
	users = uu.Repository.GetAll()
	return users
}

func (uu* UserUsecases) GetUserById (id uint) any {
	var user models.User
	user = uu.Repository.GetById(id)
	return user
}

func (uu* UserUsecases) GetOneUser (data interface{}) models.User {
	var user models.User
	user = uu.Repository.GetOne(data)
	return user
}

func (uu* UserUsecases) UpdateUser (id uint, data interface{}) models.User {
	var user models.User
	uu.Repository.Update(id, data)
	return user
}

func (uu* UserUsecases) DeleteUser (id uint) {
	uu.Repository.Delete(id)
}
