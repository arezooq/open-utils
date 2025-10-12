package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Ctx     *gin.Context
	Lang    string
	Service string
	Version string
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
