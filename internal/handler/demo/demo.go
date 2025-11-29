package demo

import (
	"fmt"
	"scaffold/errors"
	"scaffold/internal/service"
	"scaffold/response"

	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	demoSvc *service.DemoService
}

func NewDemoHandler(demoSvc *service.DemoService) *DemoHandler {
	return &DemoHandler{demoSvc: demoSvc}
}

type DemoRequest struct {
	Name string `form:"name"`
}

type DeomResponse struct {
	Reply string `json:"reply"`
}

func (demo *DemoHandler) Hello(c *gin.Context) {
	var (
		err  error
		req  DemoRequest
		resp DeomResponse
	)
	defer func() {
		if err != nil {
			response.Error(c, err)
		} else {
			response.Sucess(c, resp)
		}
	}()

	if err = c.BindQuery(&req); err != nil {
		err = errors.WithMessage(errors.BadParameters, err.Error())
		return
	}
	if req.Name == "" {
		err = errors.BadParameters
		return
	}
	resp.Reply = fmt.Sprintf("Hi, %s", req.Name)
}
