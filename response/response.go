package response

import (
	"io"
	"net/http"
	"net/url"
	"scaffold/errors"
	"scaffold/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Sucess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	if err == nil {
		Sucess(c, nil)
		return
	}

	path := c.Request.URL
	method := c.Request.Method
	contentType := c.Request.Header.Get("Content-Type")

	var form url.Values
	var body string
	switch contentType {
	case gin.MIMEJSON:
		if c.Request.Body != nil {
			content, err := io.ReadAll(c.Request.Body)
			if err == nil {
				body = string(content)
			}
		}
	case gin.MIMEPOSTForm, gin.MIMEMultipartPOSTForm:
		c.Request.ParseForm()
		form = c.Request.Form
	}

	e := &errors.Error{}
	ok := errors.As(err, e)
	if ok {
		logger.Errorf("method: %s, url: %s, form: %s, body: %s, err: %+v", method, path, form, body, e)
		c.JSON(http.StatusOK, Response{
			Status:  e.Code(),
			Message: e.Message(),
		})
	} else {
		logger.Errorf("method: %s, url: %s, form: %s, body: %s, err: %+v", method, path, form, body, e)
		c.JSON(http.StatusOK, Response{
			Status:  errors.InternalError.Code(),
			Message: errors.InternalError.Message(),
		})
	}
}
