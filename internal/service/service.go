package service

import (
	"scaffold/internal/repository"
	"scaffold/pkg/mongodb"
	"scaffold/pkg/redis"
	"scaffold/pkg/tdengine"
)

type Services struct {
	*DemoService
}

func NewServices(mongo *mongodb.Mongo, rdb *redis.Redis, taos *tdengine.Taos) *Services {
	demoRepo := repository.NewDemoRepository(mongo)

	return &Services{
		DemoService: NewDemoService(demoRepo),
	}
}
