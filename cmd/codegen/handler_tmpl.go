package main

const handlerTmpl = `
package {{.Name | toLower}}

import (
	"fmt"
	"scaffold/errors"
	"scaffold/internal/service"
	"scaffold/response"

	"github.com/gin-gonic/gin"
)

type {{.Name}}Handler struct {
	{{.Name | toLower}}Svc *service.{{.Name}}Service
}

func New{{.Name}}Handler({{.Name | toLower}}Svc *service.{{.Name}}Service) *{{.Name}}Handler {
	return &{{.Name}}Handler{
		{{.Name | toLower}}Svc: {{.Name | toLower}}Svc,
	}
}
`

const routerTmpl = `
package {{.Name | toLower}}

import "github.com/gin-gonic/gin"

func ({{.Name | toLower}} *{{.Name}}Handler) Name() string {
	return "{{.Name | toLower}}"
}

func ({{.Name | toLower}} *{{.Name}}Handler) RegisterRouter(engine *gin.RouterGroup) {
	r := engine.Group("/{{.Name | toLower}}")
	{
		// define router here
	}
}
`
