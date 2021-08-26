package usecase_test

import (
	"context"
	"errors"
	"frieda-golang-training-beginner/domain"
	mocks "frieda-golang-training-beginner/mock"
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

func TestGetPaymentCode(t *testing.T) {
	testCases := []struct {
		desc          string
		mockRepo      mocks.IPaymentCodeRepository
		expectedError error
	}{
		{
			desc: "success",
			mockRepo: func() mocks.IPaymentCodeRepository {
				repo := new(mocks.IPaymentCodeRepository)
				repo.On("GetByID", mock.Anything, mock.Anything).Return(domain.PaymentCode{Name: "Test"}, nil).Once()
				return *repo
			}(),
			expectedError: nil,
		},
		{
			desc: "nil result",
			mockRepo: func() mocks.IPaymentCodeRepository {
				repo := new(mocks.IPaymentCodeRepository)
				repo.On("GetByID", mock.Anything, mock.Anything).Return(domain.PaymentCode{}, nil).Once()
				return *repo
			}(),
			expectedError: nil,
		},
		{
			desc: "failed",
			mockRepo: func() mocks.IPaymentCodeRepository {
				repo := new(mocks.IPaymentCodeRepository)
				repo.On("GetByID", mock.Anything, mock.Anything).Return(domain.PaymentCode{}, errors.New("connection timed out")).Once()
				return *repo
			}(),
			expectedError: errors.New("connection timed out"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run(tC.desc, func(t *testing.T) {
				service := usecase.NewPaymentCodeUsecase(&tC.mockRepo, 5)
				_, err := service.Get(context.TODO(), "1")

				require.Equal(t, tC.expectedError, err)
			})
		})
	}
}

func TestExpire(t *testing.T) {
	testCases := []struct {
		desc          string
		mockRepo      mocks.IPaymentCodeRepository
		expectedError error
	}{
		{
			desc: "success",
			mockRepo: func() mocks.IPaymentCodeRepository {
				repo := new(mocks.IPaymentCodeRepository)
				repo.On("Expire", mock.Anything, mock.Anything).Return(nil).Once()
				return *repo
			}(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run(tC.desc, func(t *testing.T) {
				service := usecase.NewPaymentCodeUsecase(&tC.mockRepo, 5)
				err := service.Expire(context.TODO())
				require.Equal(t, tC.expectedError, err)
			})
		})
	}
}
