package main

const serviceTmpl = `package {{.Name | toLower}}

import "scaffold/internal/repository/{{.Name | toLower}}"

type {{.Name}}Service struct {
	{{.LowerName}}Repo {{.Name | toLower}}.{{.Name}}Repository
}

func New{{.Name}}Service({{.LowerName}}Repo {{.Name | toLower}}.{{.Name}}Repository) *{{.Name}}Service {
	return &{{.Name}}Service{
		{{.LowerName}}Repo: {{.LowerName}}Repo,
	}
}

// 业务实现方法

`
