package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajal-jayanto/go-http-template/types"
	"github.com/sajal-jayanto/go-http-template/utils"
)

type Handler struct {
	sampleRepo SampleRepo
}

func NewHandler(sampleRepo SampleRepo) *Handler {
	return &Handler{sampleRepo: sampleRepo}
}

func (handler *Handler) CreateSample(ctx *gin.Context) {	
	var body types.CreateSampleReq
	if utils.BindAndValidate(ctx, &body) {
		return 
	}

	utils.JSON(ctx, http.StatusOK, gin.H{ "body" : body.Data })
}

func (handler *Handler) Two(ctx *gin.Context) {	
	utils.JSON(ctx, http.StatusOK, gin.H{ "two" : "two" })
}