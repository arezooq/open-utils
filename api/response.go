package api

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, msg string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, code string, msg string, meta map[string]string) {
	c.JSON(status, Response{
		Success: false,
		Error: gin.H{
			"code":    code,
			"message": msg,
			"meta":    meta,
		},
	})
}
