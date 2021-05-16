package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type GetPaymentCodeResponsePayload struct {
	Status      string    `json:"status"`
	ID          uuid.UUID `json:"id"`
	PaymentCode string    `json:"payment_code"`
	Name        string    `json:"name"`
}

type CreatePaymentCodeResponsePayload struct {
	ID             uuid.UUID `json:"id"`
	PaymentCode    string    `json:"payment_code"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PaymentCode struct {
	ID             uuid.UUID `json:"id"`
	PaymentCode    string    `json:"payment_code" validate:"required"`
	Name           string    `json:"name" validate:"required"`
	Status         string    `json:"status" validate:"required"`
	ExpirationDate time.Time `json:"expiration_date" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreatePaymentCodeRequestPayload struct {
	PaymentCode    string `json:"payment_code"`
	Name           string `json:"name"`
	ExpirationDate string `json:"expiration_date"`
}

type PaymentCodeUsecase interface {
	Get(ctx context.Context, uuid uuid.UUID) (GetPaymentCodeResponsePayload, error)
	Create(ctx context.Context, request CreatePaymentCodeRequestPayload) (CreatePaymentCodeResponsePayload, error)
}

type PaymentCodeRepository interface {
	GetByID(ctx context.Context, uuid uuid.UUID) (PaymentCode, error)
	Create(ctx context.Context, paymentCode *PaymentCode) error
}
