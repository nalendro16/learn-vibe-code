package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse standardizes all JSON responses returned by the API
type APIResponse struct {
	Message string `json:"message"`
	Result  string `json:"result"`
	Data    any    `json:"data"`
}

// SendSuccess sends a standardized success response
func SendSuccess(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, APIResponse{
		Message: message,
		Result:  "ok",
		Data:    data,
	})
}

// SendError sends a standardized error response
func SendError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, APIResponse{
		Message: message,
		Result:  "error",
		Data:    nil,
	})
}

// SendCreated sends a 201 Created standardized success response
func SendCreated(c *gin.Context, message string, data any) {
	SendSuccess(c, http.StatusCreated, message, data)
}

// SendOK sends a 200 OK standardized success response
func SendOK(c *gin.Context, message string, data any) {
	SendSuccess(c, http.StatusOK, message, data)
}

// SendBadRequest sends a 400 Bad Request standardized error response
func SendBadRequest(c *gin.Context, message string) {
	SendError(c, http.StatusBadRequest, message)
}

// SendConflict sends a 409 Conflict standardized error response
func SendConflict(c *gin.Context, message string) {
	SendError(c, http.StatusConflict, message)
}

// SendInternalError sends a 500 Internal Server Error standardized error response
func SendInternalError(c *gin.Context, message string) {
	SendError(c, http.StatusInternalServerError, message)
}
