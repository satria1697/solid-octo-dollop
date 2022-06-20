package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(err interface{}) gin.H {
	return gin.H{
		"message": "Error",
		"data":    nil,
		"error":   err,
	}
}

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"message": "Success",
		"data":    data,
		"error":   nil,
	}
}
