package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendError sends a standardized error response
func SendError(c *gin.Context, status int, message string, details string) {
	c.JSON(status, ErrorResponse{
		Error: NewAPIError(status, message, details),
	})
}

// SendSuccess sends a standardized success response
func SendSuccess(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func SendOK(c *gin.Context, body any) {
	c.JSON(http.StatusOK, body)
}

func SendCreated(c *gin.Context, body any) {
	c.JSON(http.StatusCreated, body)
}

func SendNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
