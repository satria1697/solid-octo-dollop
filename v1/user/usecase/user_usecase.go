package usecase

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"three/v1/user/domain"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func (u userUseCase) GetAllUserUseCase() ([]domain.User, error) {
	res, err := u.userRepository.GetAllUserRepository()
	return res, err
}

func (u userUseCase) GetUserUseCase(id string) (domain.User, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return domain.User{}, errors.New("should-int")
	}
	res, err := u.userRepository.GetUserRepository(uint(idInt))
	return res, err
}

func (u userUseCase) DeleteUserUseCase(id string) (domain.User, error) {
	fmt.Printf("delete id: %v\no\n", id)
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return domain.User{}, errors.New("should-int")
	}

	res, err := u.userRepository.DeleteUserRepository(uint(idInt))
	return res, err
}

func (u userUseCase) UpdateUserUseCase(id string, username string, password string) (domain.User, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return domain.User{}, errors.New("should-int")
	}
	res, err := u.userRepository.UpdateUserRepository(uint(idInt), username, generatePassword(password))
	return res, err
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUseCase {
	return userUseCase{
		userRepository: userRepository,
	}
}

func generatePassword(password string) string {
	if password == "" {
		return ""
	}
	generatedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(generatedPassword)
}
