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
	var sample types.CreateSampleReq
	if utils.BindAndValidate(ctx, &sample) {
		return 
	}

	date, err := handler.sampleRepo.Create(ctx, sample)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err)
	}

	utils.JSON(ctx, http.StatusCreated, date)
}

func (handler *Handler) Two(ctx *gin.Context) {	
	utils.JSON(ctx, http.StatusOK, gin.H{ "two" : "two" })
}