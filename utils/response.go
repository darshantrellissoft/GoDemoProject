package utils

import (
	"github.com/gin-gonic/gin"
)

// APIResponse defines the structure for API responses
type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// NewResponse creates a new API response
func NewResponse(statusCode int, message string, data interface{}) APIResponse {
	return APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

// JSONResponse sends a JSON response
func JSONResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := NewResponse(statusCode, message, data)
	c.JSON(statusCode, response)
}
