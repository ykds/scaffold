package demo

import (
	"context"
	"scaffold/pkg/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	col = "demo"
)

type Demo struct {
	Name string `json:"name" db:"name"`
}

type DemoRepository interface {
	Insert(ctx context.Context, demo *Demo) error
}

type demoRepository struct {
	mgo *mongodb.Mongo
	col *mongo.Collection
}

func NewDemoRepository(mgo *mongodb.Mongo) DemoRepository {
	r := &demoRepository{
		mgo: mgo,
		col: mgo.Database.Collection(col),
	}
	return r
}

func (d *demoRepository) Insert(ctx context.Context, demo *Demo) error {
	return nil
}
