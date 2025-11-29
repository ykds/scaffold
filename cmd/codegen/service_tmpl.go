package main

const serviceTmpl = `package {{.Name | toLower}}

import "scaffold/internal/repository/{{.Name | toLower}}"

type {{.Name}}Service struct {
	{{.Name | toLower}}Repo {{.Name | toLower}}.{{.Name}}Repository
}

func New{{.Name}}Service({{.Name | toLower}}Repo {{.Name | toLower}}.{{.Name}}Repository) *{{.Name}}Service {
	return &{{.Name}}Service{
		{{.Name | toLower}}Repo: {{.Name | toLower}}Repo,
	}
}

// 业务实现方法

`
