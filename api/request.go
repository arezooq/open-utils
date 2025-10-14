package api

import (
	"fmt"
	"net/http"

	"github.com/arezooq/open-utils/errors"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Ctx     *gin.Context
	Lang    string
	Service string
	Version string
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func New(c *gin.Context, service, version string) *Request {
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "fa"
	}

	return &Request{
		Ctx:     c,
		Lang:    lang,
		Service: service,
		Version: version,
	}
}

func (r *Request) BindJSON(dst any) error {
	return r.Ctx.ShouldBindJSON(dst)
}

func (r *Request) UserValidation() (bool, error) {
	token := r.Ctx.GetHeader("Authorization")
	if token == "" {
		return false, fmt.Errorf("missing authorization header")
	}
	return true, nil
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

func FromAppError(c *gin.Context, err error, meta map[string]string) {
	if appErr, ok := err.(*errors.AppError); ok {
		c.JSON(appErr.Status, Response{
			Success: false,
			Error: gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"meta":    meta,
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error: gin.H{
			"code":    errors.ErrInternal.Code,
			"message": err.Error(),
			"meta":    meta,
		},
	})
}
