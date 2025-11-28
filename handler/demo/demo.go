package demo

import (
	"scaffold/service"

	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	demoSvc *service.DemoService
}

func NewDemoHandler(demoSvc *service.DemoService) *DemoHandler {
	return &DemoHandler{demoSvc: demoSvc}
}

func (d *DemoHandler) Hello(c *gin.Context) {
	c.JSON(200, "hello")
}

