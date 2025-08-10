package sample

import "github.com/gin-gonic/gin"

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {

	router.POST("/", handler.CreateSample)
	router.GET("/",  handler.FindAllSample)
	router.GET("/:id", handler.FindOneSample)
	router.PUT("/:id", handler.UpdateSample)
	router.DELETE("/:id", handler.RemoveSample)

}