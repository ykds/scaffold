package demo

import "github.com/gin-gonic/gin"

func (demo *DemoHandler) Name() string {
	return "demo"
}

func (demo *DemoHandler) RegisterRouter(engine *gin.RouterGroup) {
	r := engine.Group("/demo")
	{
		r.GET("", demo.Hello)
	}
}
