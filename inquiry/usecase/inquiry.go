package usecase

import (
	"context"
	"frieda-golang-training-beginner/domain"
	"github.com/jinzhu/copier"
	"time"
)

type IInquiryRepository interface {
	GetByTransactionID(ctx context.Context, transactionID string) (domain.Inquiry, error)
	Create(ctx context.Context, inquiry *domain.Inquiry) error
}

type InquiryUsecase struct {
	InquiryRepo IInquiryRepository
	ContextTimeout  time.Duration
}

func (p InquiryUsecase) Get(ctx context.Context, transactionID string) (domain.Inquiry, error) {
	var res domain.Inquiry
	paymentCode, err := p.InquiryRepo.GetByTransactionID(ctx, transactionID)
	if err != nil {
		return domain.Inquiry{}, err
	}

	err = copier.Copy(&res, &paymentCode)
	if err != nil {
		return domain.Inquiry{}, err
	}

	return res, nil
}

func (p InquiryUsecase) Create(ctx context.Context, request domain.CreateInquiryRequestPayload) (domain.Inquiry, error) {
	var inquiry domain.Inquiry
	var err error

	inquiry.PaymentCode = request.PaymentCode
	inquiry.TransactionID = request.TransactionID

	err = p.InquiryRepo.Create(ctx, &inquiry)
	if err != nil {
		return domain.Inquiry{}, err
	}
	return inquiry, nil
}

func NewInquiryUsecase(p IInquiryRepository, timeout time.Duration) InquiryUsecase {
	return InquiryUsecase{
		InquiryRepo: p,
		ContextTimeout:  timeout,
	}
}
