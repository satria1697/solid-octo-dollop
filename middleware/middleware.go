package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"three/utils"
)

func JwtMiddleware(context *gin.Context) {
	authorization := context.GetHeader("Authorization")
	if authorization == "" {
		context.JSON(http.StatusUnauthorized, utils.ErrorResponse(errors.New("not authorized")))
		context.Abort()
	} else {
		jwt := strings.Split(authorization, " ")
		if len(jwt) < 2 {
			context.JSON(http.StatusUnauthorized, utils.ErrorResponse(errors.New("not authorized")))
			context.Abort()
		} else {
			jwtToken := jwt[1]
			if jwtToken != "jwttoken" {
				context.JSON(http.StatusUnauthorized, utils.ErrorResponse(errors.New("not authorized")))
				context.Abort()
			}
		}
	}
}
