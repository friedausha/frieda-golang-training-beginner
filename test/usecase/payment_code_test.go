package test

import (
	"context"
	"errors"
	"frieda-golang-training-beginner/domain"
	"frieda-golang-training-beginner/mock"
	"frieda-golang-training-beginner/payment_code/usecase"
	mocks "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePaymentCode(t *testing.T) {
	testCases := []struct {
		desc           string
		mockRepo       mock.IPaymentCodeRepository
		expectedReturn error
	}{
		{
			desc: "success",
			mockRepo: func() mock.IPaymentCodeRepository {
				repo := new(mock.IPaymentCodeRepository)
				repo.On("Create", mocks.Anything, mocks.Anything).Return(nil).Once()
				return *repo
			}(),
			expectedReturn: nil,
		},
		{
			desc: "failed",
			mockRepo: func() mock.IPaymentCodeRepository {
				repo := new(mock.IPaymentCodeRepository)
				repo.On("Create", mocks.Anything, mocks.Anything).Return(errors.New("connection timed out")).Once()
				return *repo
			}(),
			expectedReturn: errors.New("Unknown Error"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run(tC.desc, func(t *testing.T) {
				service := usecase.NewPaymentCodeUsecase(&tC.mockRepo, 5)
				_, err := service.Create(context.TODO(), domain.CreatePaymentCodeRequestPayload{})

				require.Equal(t, tC.expectedReturn, err)
			})
		})
	}
}
