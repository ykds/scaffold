package service

import (
	demoRepo "scaffold/internal/repository/demo"
	"scaffold/internal/service/demo"
	"scaffold/pkg/mongodb"
	"scaffold/pkg/redis"
	"scaffold/pkg/tdengine"
)

type Services struct {
	*demo.DemoService
}

func NewServices(mongo *mongodb.Mongo, rdb *redis.Redis, taos *tdengine.Taos) *Services {
	demoR := demoRepo.NewDemoRepository(mongo)

	return &Services{
		DemoService: demo.NewDemoService(demoR),
	}
}
