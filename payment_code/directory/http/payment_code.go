package http

import (
	"context"
	"fmt"
	"frieda-golang-training-beginner/domain"
	"github.com/labstack/echo"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type IPaymentCodeUsecase interface {
	Get(ctx context.Context, uuid string) (domain.GetPaymentCodeResponsePayload, error)
	Create(ctx context.Context, request domain.CreatePaymentCodeRequestPayload) (domain.PaymentCode, error)
}

type PaymentCodeHandler struct {
	paymentCodeUsecase IPaymentCodeUsecase
}

func NewPaymentCodeHandler(e *echo.Echo, us IPaymentCodeUsecase) {
	handler := &PaymentCodeHandler{
		paymentCodeUsecase: us,
	}
	e.GET("/payment-codes/:id", handler.GetPaymentCode)
	e.POST("/payment-codes", handler.CreatePaymentCode)
}

func (h *PaymentCodeHandler) GetPaymentCode(c echo.Context) error {
	idP := c.Param("id")
	res, err := h.paymentCodeUsecase.Get(c.Request().Context(), idP)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error getting payment code")
	}
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
