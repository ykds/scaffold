package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"
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
	lowerName := strings.ToLower(cfg.Name)
	files := []string{
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/repository/%s/%s.go", lowerName, lowerName)),
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/service/%s/%s.go", lowerName, lowerName)),
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/%s.go", lowerName, lowerName)),
		filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/router.go", lowerName)),
	}
	for _, f := range files {
		if _, err := os.Stat(f); os.IsExist(err) {
			panic(fmt.Sprintf("%s is exists", f))
		}
	}

	// 创建必要的目录
	tmpls := []tmpl{
		{name: "repository", template: repoTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/repository/%s/%s.go", lowerName, lowerName))},
		{name: "service", template: serviceTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/service/%s/%s.go", lowerName, lowerName))},
		{name: "handler", template: handlerTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/%s.go", lowerName, lowerName))},
		{name: "handler", template: routerTmpl, path: filepath.Join(cfg.OutputPath, fmt.Sprintf("internal/handler/%s/router.go", lowerName))},
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
			Name:       cfg.Name,
			LowerName:  firstLetterToLower(cfg.Name),
			HyphenName: camelToHyphen(cfg.Name),
			SnakeName:  camelToSnake(cfg.Name),
		}
		if err := tmpl.Execute(f, data); err != nil {
			return fmt.Errorf("生成%s代码失败: %w", t.name, err)
		}

	}

	return nil
}

type tmplData struct {
	Name       string
	LowerName  string
	HyphenName string
	SnakeName  string
}

func camelToSnake(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func camelToHyphen(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}

func firstLetterToLower(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
