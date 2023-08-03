package repositories

import (
	"github.com/olumidayy/go-shawty/src/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (u* UserRepository) Create (user models.User) uint {
	u.DB.Create(&user)
	return user.ID
}

func (u* UserRepository) GetAll () []models.User {
	var users []models.User
	u.DB.Find(&users)
	return users
}

func (u* UserRepository) GetById (id uint) models.User {
	var user models.User
	u.DB.Find(&user).Where("id = ?", id)
	return user
}

func (u* UserRepository) GetOne (data interface{}) models.User {
	var user models.User
	u.DB.First(&user).Where(data)
	return user
}

func (u* UserRepository) Update (id uint, data interface{}) models.User {
	var user models.User
	u.DB.Model(&models.User{ID: id}).Updates(data)
	return user
}

func (u* UserRepository) Delete (id uint) {
	u.DB.Delete(&models.User{ID: id})
}
