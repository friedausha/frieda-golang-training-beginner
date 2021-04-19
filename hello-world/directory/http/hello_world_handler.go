package http

import (
	"frieda-golang-training-beginner/domain"
	"github.com/labstack/echo"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HelloWorldHandler struct {
	 helloWorldUCase domain.HelloWorldUsecase
}

func NewHelloWorldHandler(e *echo.Echo, us domain.HelloWorldUsecase) {
	handler := &HelloWorldHandler{
		helloWorldUCase: us,
	}
	e.GET("/hello-world", handler.GetHelloWorld)
}

func (h *HelloWorldHandler) GetHelloWorld(c echo.Context) error {
	result := h.helloWorldUCase.Get()
	return c.JSON(http.StatusOK, result)
}