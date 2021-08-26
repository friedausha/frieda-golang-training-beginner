package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"frieda-golang-training-beginner/aws"
	"frieda-golang-training-beginner/domain"
	"frieda-golang-training-beginner/inquiry/usecase"
	"github.com/google/uuid"
	"os"
	"time"
)

var QueueUrl = os.Getenv("SQS_QUEUE_URL")
type IPaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) error
}

type PaymentUsecase struct {
	PaymentRepo    IPaymentRepository
	InquiryUsecase usecase.InquiryUsecase
	SQS            aws.SQS
	ContextTimeout time.Duration
}

func (p PaymentUsecase) Create(ctx context.Context, request domain.CreatePaymentRequestPayload) (domain.Payment, error) {
	var payment domain.Payment
	var err error

	payment.PaymentCode = request.PaymentCode
	payment.TransactionID = request.TransactionID
	payment.Name = request.Name
	payment.Amount = request.Amount
	payment.TransactionID = request.TransactionID

	inquiry, err := p.InquiryUsecase.Get(ctx, payment.TransactionID)
	if err != nil {
		return domain.Payment{}, err
	}

	if inquiry.ID == uuid.Nil {
		return domain.Payment{}, fmt.Errorf("hasn't created inquiry")
	}

	paymentStr, err := json.Marshal(payment)

	err = p.PaymentRepo.Create(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	message := domain.SendRequest{
		Body:     string(paymentStr),
		QueueURL: QueueUrl,
	}

	_, err = p.SQS.Send(ctx, &message)
	if err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}

func NewPaymentUsecase(p IPaymentRepository, i usecase.InquiryUsecase, sqs aws.SQS,
	timeout time.Duration) PaymentUsecase {
	return PaymentUsecase{
		PaymentRepo:    p,
		InquiryUsecase: i,
		SQS:            sqs,
		ContextTimeout: timeout,
	}
}
