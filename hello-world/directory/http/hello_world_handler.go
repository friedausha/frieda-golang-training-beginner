package http

import (
	"frieda-golang-training-beginner/domain"
	"github.com/labstack/echo"
	"net/http"
)

type usecase interface {
	Get() domain.HelloWorldResponsePayload
}

type HelloWorldHandler struct {
	helloWorldUCase usecase
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