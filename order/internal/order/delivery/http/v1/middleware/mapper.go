package middleware

import (
	"errors"
	apperrors "github.com/IsThatASkyline/foodDelivery/order/internal/order/application/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MapErrors(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, apperrors.ErrEntityNotFound):
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, apperrors.ErrItemsMustBeMoreThanZero):
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	log.Println(err)
}
