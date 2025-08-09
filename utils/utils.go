package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, gin.H{"error": err.Error()})
}

func JSON(ctx *gin.Context, status int, payload any) {
	ctx.JSON(status, payload)
}

func BindAndValidate(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		if errors.Is(err, io.EOF) || strings.Contains(err.Error(), "EOF") {
			ctx.JSON(http.StatusBadRequest, gin.H{ "errors": []string{"Request body is required"} })
			return true
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = ValidationErrorToText(fe)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{ "errors": out })
			return true
		}
		ctx.JSON(http.StatusBadRequest, gin.H{ "errors": err.Error() })
		return true
	}
	return false
}

func GetParamInt(ctx *gin.Context, param string) (int, bool) {
	idStr := ctx.Param(param)
	if idStr == "" {
		Error(ctx, http.StatusBadRequest, fmt.Errorf("missing %s query parameter", param))
		return 0, false
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		Error(ctx, http.StatusBadRequest, fmt.Errorf("invalid %s query parameter", param))
		return 0, false
	}
	return id, true
}

func ValidationErrorToText(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters long"
	case "email":
		return fe.Field() + " must be a valid email address"
	default:
		return fe.Field() + " is invalid"
	}
}