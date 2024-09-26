package utils

import "github.com/gin-gonic/gin"

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func RespondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, APIError{
		Message: message,
		Code:    status,
	})
}
