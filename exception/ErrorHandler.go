package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"perpustakaan/helpers"
)

func NewErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			errors := err.(error)
			RespondError := gin.H{"error": errors.Error()}
			respond := helpers.ApiResponse("Server Error", http.StatusInternalServerError, "failed", RespondError)
			c.JSON(http.StatusInternalServerError, respond)
			c.Abort()
			return
		}
	}()
	c.Next()

	if len(c.Errors) > 0 {
		errors := c.Errors.Last()
		if errors.Error() == "not Found" {
			respond := helpers.ApiResponse("Data Not Found", http.StatusNotFound, "failed", errors)
			c.JSON(http.StatusNotFound, respond)
			c.Abort()
			return
		}
		respond := helpers.ApiResponse("Server Error", http.StatusInternalServerError, "failed", errors)
		c.JSON(http.StatusInternalServerError, respond)
		c.Abort()
		return
	}
}
