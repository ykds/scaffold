package main

const repoTmpl = `package {{.Name | toLower}}

import (
	"scaffold/pkg/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

// 表名
const {{.Name | toLower}}Col = "{{.Name | toLower}}"

// 表Model
type {{.Name}} struct {
}

// 接口定义
type {{.Name}}Repository interface {
}

type {{.Name | toLower}}Repository struct {
	mgo *mongodb.Mongo
	col *mongo.Collection
}

func New{{.Name}}Repository(mgo *mongodb.Mongo) {{.Name}}Repository {
	r := &{{.Name | toLower}}Repository{
		mgo: mgo,
		col: mgo.Database.Collection({{.Name | toLower}}Col),
	}
	return r
}

// 接口具体实现

`
