package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"three/middleware"
	"three/utils"
	"three/v1/user/delivery"
	"three/v1/user/domain"
)

type UserHandler struct {
	r           *gin.RouterGroup
	userUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.RouterGroup, userUseCase domain.UserUseCase) {
	handler := UserHandler{
		r:           r,
		userUseCase: userUseCase,
	}
	r.Use(middleware.JwtMiddleware)
	r.GET("/user", handler.GetAllUser)
	r.GET("/user/:id", handler.GetUser)
	r.PATCH("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)
}

func (h UserHandler) GetAllUser(c *gin.Context) {
	res, err := h.userUseCase.GetAllUserUseCase()
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(delivery.GetAllUserMapper(res)))
}

func (h UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	res, err := h.userUseCase.GetUserUseCase(id)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (h UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("delete id: %v\no\n", id)
	res, err := h.userUseCase.DeleteUserUseCase(id)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (h UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	username := c.PostForm("username")
	password := c.PostForm("password")
	res, err := h.userUseCase.UpdateUserUseCase(id, username, password)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(res))
}
