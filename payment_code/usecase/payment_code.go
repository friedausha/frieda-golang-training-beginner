package usecase

import (
	"context"
	"fmt"
	"frieda-golang-training-beginner/domain"
	"frieda-golang-training-beginner/payment_code/repository"
	"github.com/jinzhu/copier"
	"time"
)

type IPaymentCodeRepository interface {
	GetByID(ctx context.Context, uuid string) (domain.PaymentCode, error)
	Create(ctx context.Context, paymentCode *domain.PaymentCode) error
}

type PaymentCodeUsecase struct {
	PaymentCodeRepo repository.PaymentCodeRepository
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
		fmt.Println("error disini", paymentCode)
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

func NewPaymentCodeUsecase(p repository.PaymentCodeRepository, timeout time.Duration) PaymentCodeUsecase {
	return PaymentCodeUsecase{
		PaymentCodeRepo: p,
		ContextTimeout:  timeout,
	}
}
