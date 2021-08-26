package http

import (
	"context"
	"frieda-golang-training-beginner/domain"
	"github.com/labstack/echo"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type IPaymentUsecase interface {
	Create(ctx context.Context, request domain.CreatePaymentRequestPayload) (domain.Payment, error)
}

type PaymentHandler struct {
	paymentUsecase IPaymentUsecase
}

func NewPaymentHandler(e *echo.Echo, us IPaymentUsecase) {
	handler := &PaymentHandler{
		paymentUsecase: us,
	}
	e.POST("/payment", handler.CreatePayment)
}

func (h *PaymentHandler) CreatePayment(c echo.Context) error {
	var payment domain.CreatePaymentRequestPayload
	err := c.Bind(&payment)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if payment.TransactionID == "" || payment.PaymentCode == ""  || payment.Amount == 0{
		return c.JSON(http.StatusBadRequest, "missing value")
	}
	res, err := h.paymentUsecase.Create(c.Request().Context(), payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}
