package demo

import (
	"context"
	"fmt"
	"scaffold/internal/repository/demo"
)

type DemoService struct {
	demoRepo demo.DemoRepository
}

func NewDemoService(demoRepo demo.DemoRepository) *DemoService {
	return &DemoService{
		demoRepo: demoRepo,
	}
}

func (d *DemoService) Hello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}
