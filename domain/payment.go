package domain

import (
	"github.com/google/uuid"
	"time"
)

type CreatePaymentRequestPayload struct {
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`
	PaymentCode   string  `json:"payment_code"`
	TransactionID string  `json:"transaction_id"`
}

type Payment struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Amount        float64   `json:"amount"`
	PaymentCode   string    `json:"payment_code"`
	TransactionID string    `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
