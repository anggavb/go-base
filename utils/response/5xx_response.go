package utils_response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status 500 - Internal Server Error
func JSONInternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, NewResponse("Error", nil, nil, "Internal Server Error"))
}
