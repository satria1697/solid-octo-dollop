package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"three/database/postgres"
	"three/utils"
	"three/v1"
	authhandler "three/v1/auth/delivery/http"
	authrepository "three/v1/auth/repository"
	authusecase "three/v1/auth/usecase"
	userhandler "three/v1/user/delivery/http"
	userrepository "three/v1/user/repository"
	userusecase "three/v1/user/usecase"
)

func main() {
	config := utils.GetConfig()

	postgresDb, err := postgres.InitDatabase(config.Postgres)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	r := gin.Default()
	api := newApiRouter(r)
	routerv1 := v1.NewV1Router(api)

	authRepository := authrepository.NewAuthRepository(postgresDb)
	authUseCase := authusecase.NewAuthUseCase(authRepository)
	authhandler.NewAuthHandler(routerv1, authUseCase)

	userRepository := userrepository.NewUserRepository(postgresDb)
	userUseCase := userusecase.NewUserUseCase(userRepository)
	userhandler.NewUserHandler(routerv1, userUseCase)

	err = r.Run(":2032")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func newApiRouter(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/api")
}
