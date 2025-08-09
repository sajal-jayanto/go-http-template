package utils

import "github.com/gin-gonic/gin"

func Error(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, gin.H{"error": err.Error()})
}

func JSON(ctx *gin.Context, status int, payload any) {
	ctx.JSON(status, payload)
}