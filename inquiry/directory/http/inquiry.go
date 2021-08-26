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

type IInquiryUsecase interface {
	Create(ctx context.Context, request domain.CreateInquiryRequestPayload) (domain.Inquiry, error)
}

type InquiryHandler struct {
	inquiryUsecase IInquiryUsecase
}

func NewInquiryHandler(e *echo.Echo, us IInquiryUsecase) {
	handler := &InquiryHandler{
		inquiryUsecase: us,
	}
	e.POST("/inquiry", handler.CreateInquiry)
}

func (h *InquiryHandler) CreateInquiry(c echo.Context) error {
	var inquiry domain.CreateInquiryRequestPayload
	err := c.Bind(&inquiry)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if inquiry.TransactionID == "" || inquiry.PaymentCode == "" {
		return c.JSON(http.StatusBadRequest, "missing value")
	}
	res, err := h.inquiryUsecase.Create(c.Request().Context(), inquiry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}
