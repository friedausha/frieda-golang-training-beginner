package usecase_test

import (
	"context"
	"errors"
	"frieda-golang-training-beginner/domain"
	mocks"frieda-golang-training-beginner/mock"
	"frieda-golang-training-beginner/payment_code/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePaymentCode(t *testing.T) {
	testCases := []struct {
		desc           string
		mockRepo       mocks.IPaymentCodeRepository
		expectedReturn error
	}{
		{
			desc: "success",
			mockRepo: func() mocks.IPaymentCodeRepository {
				repo := new(mocks.IPaymentCodeRepository)
				repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
				return *repo
			}(),
			expectedReturn: nil,
		},
		{
			desc: "failed",
			mockRepo: func() mocks.IPaymentCodeRepository {
				repo := new(mocks.IPaymentCodeRepository)
				repo.On("Create", mock.Anything, mock.Anything).Return(errors.New("connection timed out")).Once()
				return *repo
			}(),
			expectedReturn: errors.New("connection timed out"),
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
