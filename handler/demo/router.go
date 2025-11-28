package demo

import "github.com/gin-gonic/gin"

func (d *DemoHandler) Name() string {
	return "demo"
}

func (d *DemoHandler) RegisterRouter(engine *gin.RouterGroup) {
	r := engine.Group("/demo")
	{
		r.GET("", d.Hello)
	}
}
