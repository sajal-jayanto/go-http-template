package sample

import (
	"fmt"
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

func (handler *Handler) FindAllSample(ctx *gin.Context) {	
	samples, err := handler.sampleRepo.FindAll(ctx)
	if err != nil{
		fmt.Println(samples)
		utils.Error(ctx, http.StatusInternalServerError, err)
	}

	utils.JSON(ctx, http.StatusOK, samples)
}