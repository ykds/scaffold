package main

const handlerTmpl = `package {{.SnakeName}}

import (
	"scaffold/internal/service/{{.SnakeName}}"
)

type {{.Name}}Handler struct {
	{{.LowerName}}Svc *{{.SnakeName}}.{{.Name}}Service
}

func New{{.Name}}Handler({{.LowerName}}Svc *{{.SnakeName}}.{{.Name}}Service) *{{.Name}}Handler {
	return &{{.Name}}Handler{
		{{.LowerName}}Svc: {{.LowerName}}Svc,
	}
}

// 控制层代码实现

`

const routerTmpl = `package {{.SnakeName}}

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
