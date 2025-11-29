package service

import "scaffold/internal/repository"

type DemoService struct {
	demoRepo repository.DemoRepository
}

func NewDemoService(demoRepo repository.DemoRepository) *DemoService {
	return &DemoService{
		demoRepo: demoRepo,
	}
}
