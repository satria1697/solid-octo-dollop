package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"three/database/postgres"
	"three/utils"
	"three/v1"
	authhandler "three/v1/auth/delivery/http"
	authrepository "three/v1/auth/repository"
	authusecase "three/v1/auth/usecase"
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
	authhandler.NewUserHandler(routerv1, authUseCase)
	r.Run(":2032")
}

func newApiRouter(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/api")
}
