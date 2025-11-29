package main

const serviceTmpl = `package service

import "scaffold/internal/repository"

type {{.Name}}Service struct {
	{{.Name | toLower}}Repo repository.{{.Name}}Repository
}

func New{{.Name}}Service({{.Name | toLower}}Repo repository.{{.Name}}Repository) *{{.Name}}Service {
	return &{{.Name}}Service{
		{{.Name | toLower}}Repo: {{.Name | toLower}}Repo,
	}
}

`