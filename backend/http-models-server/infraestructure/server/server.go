package server

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	return gin.Default()
}
