package handler

import (
	"reflect"
	"scaffold/internal/handler/demo"
	"scaffold/internal/service"

	"github.com/gin-gonic/gin"
)

var handlers = make(map[string]Handler)

type Handler interface {
	Name() string
	RegisterRouter(engine *gin.RouterGroup)
}

type Handlers struct {
	*demo.DemoHandler
}

func InitHandlers(s *service.Services) {
	demoHandler := demo.NewDemoHandler(s.DemoService)
	h := Handlers{
		DemoHandler: demoHandler,
	}

	registerHandler(h)
}

func registerHandler(h Handlers) {
	v := reflect.ValueOf(h)
	for i := range v.NumField() {
		field := v.Field(i)
		h := field.Interface().(Handler)
		if _, ok := handlers[h.Name()]; !ok {
			handlers[h.Name()] = h
		}
	}
}

func RegisterRouter(engine *gin.Engine) {
	group := engine.Group("")
	for _, h := range handlers {
		h.RegisterRouter(group)
	}
}
