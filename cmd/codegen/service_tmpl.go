package main

const serviceTmpl = `package {{.SnakeName}}

import "scaffold/internal/repository/{{.SnakeName}}"

type {{.Name}}Service struct {
	{{.LowerName}}Repo {{.SnakeName}}.{{.Name}}Repository
}

func New{{.Name}}Service({{.LowerName}}Repo {{.SnakeName}}.{{.Name}}Repository) *{{.Name}}Service {
	return &{{.Name}}Service{
		{{.LowerName}}Repo: {{.LowerName}}Repo,
	}
}

// 业务实现方法

`
