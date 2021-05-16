package http

import (
	"frieda-golang-training-beginner/domain"
	"github.com/labstack/echo"
	"net/http"
)

type HealthHandler struct {
	healthUsecase domain.HealthUsecase
}

func NewHealthHandler(e *echo.Echo, us domain.HealthUsecase) {
	handler := &HealthHandler{
		healthUsecase: us,
	}
	e.GET("/health", handler.GetHealth)
}

func (h *HealthHandler) GetHealth(c echo.Context) error {
	result := h.healthUsecase.Get()
	return c.JSON(http.StatusOK, result)
}
