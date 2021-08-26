package usecase_test

import (
	"context"
	"fmt"
	"frieda-golang-training-beginner/domain"
	usecase2 "frieda-golang-training-beginner/inquiry/usecase"
	"frieda-golang-training-beginner/payment/usecase"
	mocks "frieda-golang-training-beginner/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePayment(t *testing.T) {
	testCases := []struct {
		desc           string
		req            domain.CreatePaymentRequestPayload
		mockRepo       mocks.IPaymentRepository
		InquiryUC      mocks.IInquiryRepository
		expectedReturn error
	}{
		{
			desc: "success",
			req: domain.CreatePaymentRequestPayload{TransactionID: "1", Name: "a"},
			mockRepo: func() mocks.IPaymentRepository {
				repo := new(mocks.IPaymentRepository)
				repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
				return *repo
			}(),
			InquiryUC: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("GetByTransactionID", mock.Anything, mock.Anything).
					Return(domain.Inquiry{ID: uuid.New()}, nil).Once()
				return *repo
			}(),
			expectedReturn: nil,
		},
		{
			desc: "failed-no-inquiry",
			req: domain.CreatePaymentRequestPayload{TransactionID: "1", Name: "a"},
			mockRepo: func() mocks.IPaymentRepository {
				repo := new(mocks.IPaymentRepository)
				repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
				return *repo
			}(),
			InquiryUC: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("GetByTransactionID", mock.Anything, mock.Anything).
					Return(domain.Inquiry{}, nil).Once()
				return *repo
			}(),
			expectedReturn: fmt.Errorf("hasn't created inquiry"),
		},
		{
			desc: "failed",
			req: domain.CreatePaymentRequestPayload{TransactionID: "1", Name: "a"},
			mockRepo: func() mocks.IPaymentRepository {
				repo := new(mocks.IPaymentRepository)
				repo.On("Create", mock.Anything, mock.Anything).Return(fmt.Errorf("500")).Once()
				return *repo
			}(),
			InquiryUC: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("GetByTransactionID", mock.Anything, mock.Anything).
					Return(domain.Inquiry{ID: uuid.New()}, nil).Once()
				return *repo
			}(),
			expectedReturn: fmt.Errorf("500"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run(tC.desc, func(t *testing.T) {
				inquiryUc := usecase2.NewInquiryUsecase(&tC.InquiryUC, 5)
				service := usecase.NewPaymentUsecase(&tC.mockRepo, inquiryUc, 5)
				_, err := service.Create(context.TODO(), tC.req)

				require.Equal(t, tC.expectedReturn, err)
			})
		})
	}
}