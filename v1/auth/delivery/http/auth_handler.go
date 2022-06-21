package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"three/utils"
	"three/v1/auth/domain"
)

type UserHandler struct {
	r           *gin.RouterGroup
	authUseCase domain.AuthUseCase
}

func NewAuthHandler(r *gin.RouterGroup, authUseCase domain.AuthUseCase) {
	handler := UserHandler{
		authUseCase: authUseCase,
	}
	authR := r.Group("/auth")
	authR.POST("/login", handler.Login)
	authR.POST("/register", handler.Register)
}

func (h UserHandler) Login(c *gin.Context) {
	var request domain.AuthLoginRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	res, err := h.authUseCase.LoginUseCase(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(res))
	return
}

func (h UserHandler) Register(c *gin.Context) {
	var request domain.AuthLoginRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	res, err := h.authUseCase.RegisterUseCase(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(res))
	return
}
