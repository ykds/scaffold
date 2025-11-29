package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Config struct {
	Name       string
	OutputPath string
}

func Generate(cfg Config) error {
	if err := generate(cfg); err != nil {
		panic(err)
	}

	return nil
}

type tmpl struct {
	name     string
	template string
	path     string
}

func generate(cfg Config) error {
	files := []string{
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/repository/%s/%s.go", strings.ToLower(cfg.Name), strings.ToLower(cfg.Name))),
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/service/%s/%s.go", strings.ToLower(cfg.Name), strings.ToLower(cfg.Name))),
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/%s.go", strings.ToLower(cfg.Name), strings.ToLower(cfg.Name))),
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/router.go", strings.ToLower(cfg.Name))),
	}
	for _, f := range files {
		if _, err := os.Stat(f); os.IsExist(err) {
			panic(fmt.Sprintf("%s is exists", f))
		}
	}

	// 创建必要的目录
	tmpls := []tmpl{
		{name: "repository", template: repoTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/repository/%s/%s.go", strings.ToLower(cfg.Name), strings.ToLower(cfg.Name)))},
		{name: "service", template: serviceTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/service/%s/%s.go", strings.ToLower(cfg.Name), strings.ToLower(cfg.Name)))},
		{name: "handler", template: handlerTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/%s.go", strings.ToLower(cfg.Name), strings.ToLower(cfg.Name)))},
		{name: "handler", template: routerTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/router.go", strings.ToLower(cfg.Name)))},
	}
	for _, f := range tmpls {
		if _, err := os.Stat(f.path); os.IsExist(err) {
			panic(fmt.Sprintf("%s is exists", f))
		}
		dir := filepath.Dir(f.path)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("创建%s目录失败: %w", dir, err)
		}
	}

	for _, t := range tmpls {
		f, err := os.Create(t.path)
		if err != nil {
			return fmt.Errorf("创建%s文件失败: %w", t.path, err)
		}

		funcMap := template.FuncMap{
			"toLower": strings.ToLower,
		}

		tmpl, err := template.New(t.name).
			Funcs(funcMap).
			Parse(t.template)
		if err != nil {
			return fmt.Errorf("解析%s模板失败: %w", t.path, err)
		}

		data := tmplData{
			Name: cfg.Name,
		}
		if err := tmpl.Execute(f, data); err != nil {
			return fmt.Errorf("生成%s代码失败: %w", t.name, err)
		}

	}

	return nil
}

type tmplData struct {
	Name string
}
