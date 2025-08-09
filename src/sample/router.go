package sample

import "github.com/gin-gonic/gin"

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {
	
	router.POST("/",   handler.CreateSample)
	router.GET("/two", handler.Two)
	
}