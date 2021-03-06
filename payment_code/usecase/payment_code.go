package usecase

import (
	"context"
	"frieda-golang-training-beginner/domain"
	"github.com/jinzhu/copier"
	"time"
)

type IPaymentCodeRepository interface {
	GetByID(ctx context.Context, uuid string) (domain.PaymentCode, error)
	Create(ctx context.Context, paymentCode *domain.PaymentCode) error
	Expire(ctx context.Context) error
}

type PaymentCodeUsecase struct {
	PaymentCodeRepo IPaymentCodeRepository
	ContextTimeout  time.Duration
}

func (p PaymentCodeUsecase) Get(ctx context.Context, uuid string) (domain.GetPaymentCodeResponsePayload, error) {
	var res domain.GetPaymentCodeResponsePayload
	paymentCode, err := p.PaymentCodeRepo.GetByID(ctx, uuid)
	if err != nil {
		return domain.GetPaymentCodeResponsePayload{}, err
	}

	err = copier.Copy(&res, &paymentCode)
	if err != nil {
		return domain.GetPaymentCodeResponsePayload{}, err
	}

	return res, nil
}

func (p PaymentCodeUsecase) Create(ctx context.Context, request domain.CreatePaymentCodeRequestPayload) (domain.PaymentCode, error) {
	var paymentCode domain.PaymentCode
	var err error

	paymentCode.PaymentCode = request.PaymentCode
	paymentCode.Name = request.Name
	paymentCode.Status = request.Status
	if request.ExpirationDateString == "" {
		paymentCode.ExpirationDate = time.Now().AddDate(51, 0, 0).UTC()
	} else {
		paymentCode.ExpirationDate, err = time.Parse(time.RFC3339, request.ExpirationDateString)
	}

	err = p.PaymentCodeRepo.Create(ctx, &paymentCode)
	if err != nil {
		return domain.PaymentCode{}, err
	}
	return paymentCode, nil
}

func (p PaymentCodeUsecase) Expire(ctx context.Context) error {
	err := p.PaymentCodeRepo.Expire(ctx)
	if err != nil {
		return err
	}
	return err
}

func NewPaymentCodeUsecase(p IPaymentCodeRepository, timeout time.Duration) PaymentCodeUsecase {
	return PaymentCodeUsecase{
		PaymentCodeRepo: p,
		ContextTimeout:  timeout,
	}
}
