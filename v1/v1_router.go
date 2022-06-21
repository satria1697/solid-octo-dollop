package v1

import "github.com/gin-gonic/gin"

func NewV1Router(r *gin.RouterGroup) *gin.RouterGroup {
	return r.Group("/v1")
}
