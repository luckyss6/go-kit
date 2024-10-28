package server

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	return app
}
