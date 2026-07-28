package utils_response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status 200 - OK
func JSONSuccess(ctx *gin.Context, data, meta any, message string) {
	ctx.JSON(http.StatusOK, NewResponse(message, data, meta, nil))
}

// Status 201 - Created
func JSONCreated(ctx *gin.Context, data, meta any, message string) {
	ctx.JSON(http.StatusCreated, NewResponse(message, data, meta, nil))
}

// Status 204 - No Content
func JSONNoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
