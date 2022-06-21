package repository

import (
	"fmt"
	"gorm.io/gorm"
	"three/v1/user/domain"
)

type userRepository struct {
	postgresDb *gorm.DB
}

func (u userRepository) GetAllUserRepository() ([]domain.User, error) {
	var users []domain.User

	resDb := u.postgresDb.Find(&users)
	if resDb.Error != nil {
		return users, resDb.Error
	}
	return users, nil
}

func (u userRepository) GetUserRepository(id uint) (domain.User, error) {
	var user domain.User
	resDb := u.postgresDb.First(&user, id)
	if resDb.Error != nil {
		return user, resDb.Error
	}
	return user, nil
}

func (u userRepository) DeleteUserRepository(id uint) (domain.User, error) {
	var user domain.User
	fmt.Printf("%v\n", id)
	resDb := u.postgresDb.Delete(&user, id)
	if resDb.Error != nil {
		return user, resDb.Error
	}
	return user, nil
}

func (u userRepository) UpdateUserRepository(id uint, username string, password string) (domain.User, error) {
	user := domain.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	resDb := u.postgresDb.Model(&user).Updates(domain.User{
		Username: username,
		Password: password,
	})
	if resDb.Error != nil {
		return user, resDb.Error
	}
	return user, nil
}

func NewUserRepository(postgresDb *gorm.DB) domain.UserRepository {
	return userRepository{
		postgresDb: postgresDb,
	}
}
