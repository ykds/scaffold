package main

const handlerTmpl = `package {{.Name | toLower}}

import (
	"scaffold/internal/service/{{.Name | toLower}}"
)

type {{.Name}}Handler struct {
	{{.LowerName}}Svc *{{.Name | toLower}}.{{.Name}}Service
}

func New{{.Name}}Handler({{.LowerName}}Svc *{{.Name | toLower}}.{{.Name}}Service) *{{.Name}}Handler {
	return &{{.Name}}Handler{
		{{.LowerName}}Svc: {{.LowerName}}Svc,
	}
}

// 控制层代码实现

`

const routerTmpl = `package {{.Name | toLower}}

import "github.com/gin-gonic/gin"

func ({{.LowerName}} *{{.Name}}Handler) Name() string {
	return "{{.LowerName}}"
}

func ({{.LowerName}} *{{.Name}}Handler) RegisterRouter(engine *gin.RouterGroup) {
	_ = engine.Group("/{{.HyphenName}}")
	{
		// define router here
	}
}
`
