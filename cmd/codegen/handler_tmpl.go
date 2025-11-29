package main

const handlerTmpl = `package {{.Name | toLower}}

import (
	"scaffold/internal/service/{{.Name | toLower}}"
)

type {{.Name}}Handler struct {
	{{.Name | toLower}}Svc *{{.Name | toLower}}.{{.Name}}Service
}

func New{{.Name}}Handler({{.Name | toLower}}Svc *{{.Name | toLower}}.{{.Name}}Service) *{{.Name}}Handler {
	return &{{.Name}}Handler{
		{{.Name | toLower}}Svc: {{.Name | toLower}}Svc,
	}
}

// 控制层代码实现

`

const routerTmpl = `package {{.Name | toLower}}

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
