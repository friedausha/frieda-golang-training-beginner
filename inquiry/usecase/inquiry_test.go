package usecase_test

import (
	"context"
	"errors"
	"frieda-golang-training-beginner/domain"
	mocks "frieda-golang-training-beginner/mock"
	"frieda-golang-training-beginner/inquiry/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateInquiry(t *testing.T) {
	testCases := []struct {
		desc           string
		mockRepo       mocks.IInquiryRepository
		expectedReturn error
	}{
		{
			desc: "success",
			mockRepo: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
				return *repo
			}(),
			expectedReturn: nil,
		},
		{
			desc: "failed",
			mockRepo: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("Create", mock.Anything, mock.Anything).Return(errors.New("connection timed out")).Once()
				return *repo
			}(),
			expectedReturn: errors.New("connection timed out"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run(tC.desc, func(t *testing.T) {
				service := usecase.NewInquiryUsecase(&tC.mockRepo, 5)
				_, err := service.Create(context.TODO(), domain.CreateInquiryRequestPayload{})

				require.Equal(t, tC.expectedReturn, err)
			})
		})
	}
}

func TestGetInquiry(t *testing.T) {
	testCases := []struct {
		desc          string
		mockRepo      mocks.IInquiryRepository
		expectedError error
	}{
		{
			desc: "success",
			mockRepo: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("GetByTransactionID", mock.Anything, mock.Anything).Return(domain.Inquiry{TransactionID: "Test"}, nil).Once()
				return *repo
			}(),
			expectedError: nil,
		},
		{
			desc: "nil result",
			mockRepo: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("GetByTransactionID", mock.Anything, mock.Anything).Return(domain.Inquiry{}, nil).Once()
				return *repo
			}(),
			expectedError: nil,
		},
		{
			desc: "failed",
			mockRepo: func() mocks.IInquiryRepository {
				repo := new(mocks.IInquiryRepository)
				repo.On("GetByTransactionID", mock.Anything, mock.Anything).Return(domain.Inquiry{}, errors.New("connection timed out")).Once()
				return *repo
			}(),
			expectedError: errors.New("connection timed out"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run(tC.desc, func(t *testing.T) {
				service := usecase.NewInquiryUsecase(&tC.mockRepo, 5)
				_, err := service.Get(context.TODO(), "1")

				require.Equal(t, tC.expectedError, err)
			})
		})
	}
}
