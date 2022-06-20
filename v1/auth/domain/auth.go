package domain

import "three/v1/user/domain"

type AuthLoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthRegisterRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthUseCase interface {
	LoginUseCase(username string, password string) (string, error)
	RegisterUseCase(username string, password string) (domain.User, error)
}

type AuthRepository interface {
	LoginRepository(username string, passwowrd string) (string, error)
	RegisterRepository(username string, password string) (domain.User, error)
}
