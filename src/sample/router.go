package sample

import "github.com/gin-gonic/gin"

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/one", handler.One)
	router.GET("/two", handler.Two)
}