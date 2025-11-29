package demo

import "scaffold/internal/repository/demo"

type DemoService struct {
	demoRepo demo.DemoRepository
}

func NewDemoService(demoRepo demo.DemoRepository) *DemoService {
	return &DemoService{
		demoRepo: demoRepo,
	}
}
