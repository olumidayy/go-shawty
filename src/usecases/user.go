package usecases

import (
	"github.com/olumidayy/go-shawty/src/domain/models"
)

type UserRepository interface {
	Create(user models.User) (*models.User, error)
	GetAll() (*[]models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetOne(data interface{}) (*models.User, error)
	Update(id uint, user *models.User) (*models.User, error)
	Delete(id uint) error
}

type UserUsecases struct {
	Repository UserRepository
}

func NewUserUsecases(r UserRepository) *UserUsecases {
	return &UserUsecases{
		Repository: r,
	}
}

func (uu* UserUsecases) CreateNewUser (user models.User) (*models.User, error) {
	return uu.Repository.Create(user)
}

func (uu* UserUsecases) GetAllUsers () (*[]models.User, error) {
	return uu.Repository.GetAll()
}

func (uu* UserUsecases) GetUserByID (id uint) (*models.User, error) {
	return uu.Repository.GetByID(id)
}

func (uu* UserUsecases) GetUserByEmail (id uint) (*models.User, error) {
	return uu.Repository.GetByID(id)
}

func (uu* UserUsecases) GetOneUser (data *models.User) (*models.User, error) {
	return uu.Repository.GetOne(data)
}

func (uu* UserUsecases) UpdateUser (id uint, data *models.User) (*models.User, error) {
	return uu.Repository.Update(id, data)
}

func (uu* UserUsecases) DeleteUser (id uint) error {
	return uu.Repository.Delete(id)
}
