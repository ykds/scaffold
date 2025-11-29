package main

const repoTmpl = `package {{.SnakeName}}

import (
	"scaffold/pkg/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

// 表名
const {{.LowerName}}Col = "{{.SnakeName}}"

// 表Model
type {{.Name}} struct {
}

// 接口定义
type {{.Name}}Repository interface {
}

type {{.LowerName}}Repository struct {
	mgo *mongodb.Mongo
	col *mongo.Collection
}

func New{{.Name}}Repository(mgo *mongodb.Mongo) {{.Name}}Repository {
	r := &{{.LowerName}}Repository{
		mgo: mgo,
		col: mgo.Database.Collection({{.LowerName}}Col),
	}
	return r
}

// 接口具体实现

`
