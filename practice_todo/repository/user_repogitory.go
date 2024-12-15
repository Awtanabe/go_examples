package repository

import (
	"practice_todo/models"

	"gorm.io/gorm"
)

type IUserRepogitory interface {
	// 型を渡してる？
	GetUser(userId uint, user *models.User) error
	GetUserByEmail(user *models.User, email string) error
	CreateUser(user *models.User) error
}


type userRepogitory struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepogitory {
	return &userRepogitory{db}
}

func (ur *userRepogitory) GetUser(userId uint, user *models.User) error {
	if err := ur.db.Model(&models.User{}).
			Preload("Todos").
			Where("id = ?", 1).
			First(user).Error; err != nil {
			return err
	}
	return nil
}

func (ur *userRepogitory) GetUserByEmail(user *models.User, email string) error {
	// db.whereはgormの機能
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepogitory) CreateUser(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}