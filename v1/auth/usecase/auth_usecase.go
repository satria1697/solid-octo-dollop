package usecase

import (
	"three/utils"
	"three/v1/auth/domain"
	userdomain "three/v1/user/domain"
)

type authUseCase struct {
	authRepository domain.AuthRepository
}

func (a authUseCase) RegisterUseCase(username string, password string) (userdomain.User, error) {
	res, err := a.authRepository.RegisterRepository(username, password)
	return res, err
}

func (a authUseCase) LoginUseCase(username string, password string) (string, error) {
	res, err := a.authRepository.LoginRepository(username, password)
	if err != nil {
		return "", err
	}
	return utils.CreateToken(res.Username)
}

func NewAuthUseCase(authRepository domain.AuthRepository) domain.AuthUseCase {
	return authUseCase{
		authRepository: authRepository,
	}
}
