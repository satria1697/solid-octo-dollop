package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"three/v1/auth/domain"
	userdomain "three/v1/user/domain"
)

type authRepository struct {
	postgresDb *gorm.DB
}

func (a authRepository) RegisterRepository(username string, password string) (userdomain.User, error) {
	var user userdomain.User
	resDb := a.postgresDb.Find(&user, "username = ?", username)
	if resDb.RowsAffected > 0 {
		return userdomain.User{}, errors.New("username_exist")
	}
	generatePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return userdomain.User{}, err
	}
	resDb = a.postgresDb.Create(&userdomain.User{
		Username: username,
		Password: string(generatePassword),
	})
	if resDb.Error != nil {
		return userdomain.User{}, resDb.Error
	}
	return userdomain.User{}, nil
}

func (a authRepository) LoginRepository(username string, password string) (userdomain.User, error) {
	var user userdomain.User
	resDb := a.postgresDb.Find(&user, "username = ?", username)
	if resDb.RowsAffected == 0 {
		return userdomain.User{}, errors.New("username_password_false")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return userdomain.User{}, errors.New("username_password_false")
	}
	return userdomain.User{
		Username: username,
	}, nil
}

func NewAuthRepository(postgresDb *gorm.DB) domain.AuthRepository {
	return authRepository{
		postgresDb: postgresDb,
	}
}
