package utils_response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status 400 - Bad Request
func JSONBadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, NewResponse("Invalid Request Payload", nil, nil, "Bad Request"))
}

// Status 400 - Bad Request with custom message
func JSONBadRequestWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, NewResponse(message, nil, nil, "Bad Request"))
}

// Status 401 - Unauthorized
func JSONUnauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, NewResponse(message, nil, nil, "Unauthorized"))
}

// Status 409 - Conflict
func JSONConflict(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusConflict, NewResponse(message, nil, nil, "Conflict"))
}

// Status 404 - Not Found
func JSONNotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, NewResponse(message, nil, nil, "Not Found"))
}

// Status 422 - Unprocessable Entity
func JSONUnprocessableEntity(ctx *gin.Context, errors map[string]string) {
	ctx.JSON(http.StatusUnprocessableEntity, NewResponse("Unprocessable Entity", nil, nil, errors))
}
