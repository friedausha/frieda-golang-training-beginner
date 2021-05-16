package http

import (
	"frieda-golang-training-beginner/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PaymentCodeHandler struct {
	paymentCodeUsecase domain.PaymentCodeUsecase
}

func NewPaymentCodeHandler(e *echo.Echo, us domain.PaymentCodeUsecase) {
	handler := &PaymentCodeHandler{
		paymentCodeUsecase: us,
	}
	e.GET("/payment-codes/:id", handler.GetPaymentCode)
	e.POST("/payment-codes", handler.CreatePaymentCode)
}

func (h *PaymentCodeHandler) GetPaymentCode(c echo.Context) error {
	idP, err := uuid.FromBytes([]byte(c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error converting id")
	}
	res, err := h.paymentCodeUsecase.Get(c.Request().Context(), idP)
	return c.JSON(http.StatusOK, res)
}

func (h *PaymentCodeHandler) CreatePaymentCode(c echo.Context) error {
	var paymentCode domain.CreatePaymentCodeRequestPayload
	err := c.Bind(&paymentCode)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	res, err := h.paymentCodeUsecase.Create(c.Request().Context(), paymentCode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}
