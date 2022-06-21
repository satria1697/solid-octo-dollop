package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{
		"message": "Error",
		"data":    nil,
		"error":   err.Error(),
	}
}

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"message": "Success",
		"data":    data,
		"error":   nil,
	}
}
