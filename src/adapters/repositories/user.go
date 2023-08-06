package repositories

import (
	"fmt"

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

func (u* UserRepository) Create (user models.User) (*models.User, error) {
	err := u.DB.Create(&user).Error;
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u* UserRepository) GetAll () (*[]models.User, error) {
	var users *[]models.User
	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u* UserRepository) GetByID (id uint) (*models.User, error) {
	var user models.User
	err := u.DB.Where(&models.User{ID: id}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u* UserRepository) GetByEmail (email string) (*models.User, error) {
	var user models.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u* UserRepository) GetOne (data interface{}) (*models.User, error) {
	var user models.User
	err := u.DB.Where(data).First(&user).Error
	if err != nil {
		fmt.Println("i'm here")
		return nil, err
	}
	fmt.Println(user)
	return &user, nil
}

func (u* UserRepository) Update (id uint, user *models.User) (*models.User, error) {
	err := u.DB.Model(&models.User{ID: id}).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u* UserRepository) Delete (id uint) error {
	return u.DB.Delete(&models.User{ID: id}).Error
}
