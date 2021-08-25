package domain

import (
	"github.com/google/uuid"
	"time"
)

type CreateInquiryRequestPayload struct {
	PaymentCode   string `json:"payment_code"`
	TransactionID string `json:"transaction_id"`
}

type Inquiry struct {
	ID             uuid.UUID `json:"id"`
	PaymentCode    string    `json:"payment_code" validate:"required"`
	TransactionID  string    `json:"transaction_id" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
