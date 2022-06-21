package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type GetAllUserRespose struct {
	ID       uint   `json:"ID"`
	Username string `json:"username"`
}

type UserUseCase interface {
	GetAllUserUseCase() ([]User, error)
	GetUserUseCase(id string) (User, error)
	DeleteUserUseCase(id string) (User, error)
	UpdateUserUseCase(id string, username string, password string) (User, error)
}

type UserRepository interface {
	GetAllUserRepository() ([]User, error)
	GetUserRepository(id uint) (User, error)
	DeleteUserRepository(id uint) (User, error)
	UpdateUserRepository(id uint, username string, password string) (User, error)
}
