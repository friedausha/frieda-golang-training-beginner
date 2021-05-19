package domain

import (
	"github.com/google/uuid"
	"time"
)

type GetPaymentCodeResponsePayload struct {
	Status      string    `json:"status"`
	ID          uuid.UUID `json:"id"`
	PaymentCode string    `json:"payment_code"`
	Name        string    `json:"name"`
}

type PaymentCode struct {
	ID             uuid.UUID `json:"id"`
	PaymentCode    string    `json:"payment_code" validate:"required"`
	Name           string    `json:"name" validate:"required"`
	Status         string    `json:"status" validate:"required"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreatePaymentCodeRequestPayload struct {
	PaymentCode          string `json:"payment_code"`
	Name                 string `json:"name"`
	ExpirationDateString string `json:"expiration_date"`
	Status               string `json:"status"`
}
