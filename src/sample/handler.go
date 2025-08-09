package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajal-jayanto/go-http-template/utils"
)

type Handler struct {
	sampleRepo SampleRepo
}

func NewHandler(sampleRepo SampleRepo) *Handler {
	return &Handler{sampleRepo: sampleRepo}
}

func (handler *Handler) One(ctx *gin.Context) {	
	utils.JSON(ctx, http.StatusOK, gin.H{ "one" : "one" })
}

func (handler *Handler) Two(ctx *gin.Context) {	
	utils.JSON(ctx, http.StatusOK, gin.H{ "two" : "two" })
}