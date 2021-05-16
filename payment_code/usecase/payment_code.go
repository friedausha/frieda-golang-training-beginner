package usecase

import (
	"context"
	"frieda-golang-training-beginner/domain"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"time"
)

type paymentCodeUsecase struct {
	paymentCodeRepo domain.PaymentCodeRepository
	contextTimeout  time.Duration
}

func (p paymentCodeUsecase) Get(ctx context.Context, uuid uuid.UUID) (domain.GetPaymentCodeResponsePayload, error) {
	var res domain.GetPaymentCodeResponsePayload
	paymentCode, err := p.paymentCodeRepo.GetByID(ctx, uuid)
	if err != nil {
		return domain.GetPaymentCodeResponsePayload{}, err
	}

	err = copier.Copy(&res, &paymentCode)
	if err != nil {
		return domain.GetPaymentCodeResponsePayload{}, err
	}

	return res, nil
}

func (p paymentCodeUsecase) Create(ctx context.Context, request domain.CreatePaymentCodeRequestPayload) (domain.CreatePaymentCodeResponsePayload, error) {
	var paymentCode *domain.PaymentCode
	var resp domain.CreatePaymentCodeResponsePayload
	err := copier.Copy(request, paymentCode)
	if err != nil {
		return domain.CreatePaymentCodeResponsePayload{}, err
	}

	err = p.paymentCodeRepo.Create(ctx, paymentCode)
	err = copier.Copy(&paymentCode, &resp)
	if err != nil {
		return domain.CreatePaymentCodeResponsePayload{}, err
	}
	return resp, nil
}

func NewPaymentCodeUsecase(p domain.PaymentCodeRepository, timeout time.Duration) domain.PaymentCodeUsecase {
	return &paymentCodeUsecase{
		paymentCodeRepo: p,
		contextTimeout:  timeout,
	}
}
