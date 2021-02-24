package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ok
func Ok(obj interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, obj)
}

// BadRequestError
func BadRequestError(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, msg)
}

// InternalServerError
func InternalServerError(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, msg)
}

// Unauthorized
func Unauthorized(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, msg)
}

// Forbidden
func Forbidden(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, msg)
}

// NotFound
func NotFound(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, msg)
}

// NoContent
func NoContent(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, msg)
}
