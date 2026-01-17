package response

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}

// Error sends an error response
func Error(c *gin.Context, statusCode int, message string, err any) {
	var errorDetail any
	
	// If error is nil, don't include it in the response
	if err != nil {
		switch e := err.(type) {
		case error:
			errorDetail = e.Error()
		default:
			errorDetail = e
		}
	}
	
	c.JSON(statusCode, ErrorResponse{
		Success: false,
		Message: message,
		Error:   errorDetail,
	})
}