package http

import (
	"frieda-golang-training-beginner/domain"
	"github.com/labstack/echo"
	"net/http"
)

type IHealthUsecase interface {
	Get() domain.HealthResponsePayload
}

type HealthHandler struct {
	healthUsecase IHealthUsecase
}

func NewHealthHandler(e *echo.Echo, us IHealthUsecase) {
	handler := &HealthHandler{
		healthUsecase: us,
	}
	e.GET("/health", handler.GetHealth)
}


func (h *HealthHandler) GetHealth(c echo.Context) error {
	result := h.healthUsecase.Get()
	return c.JSON(http.StatusOK, result)
}