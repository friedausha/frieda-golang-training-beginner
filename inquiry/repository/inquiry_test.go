package repository_test

import (
	"context"
	"frieda-golang-training-beginner/domain"
	"frieda-golang-training-beginner/inquiry/repository"
	"frieda-golang-training-beginner/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type inquiryTestSuite struct {
	util.Suite
}

func TestSuiteInquiry(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test Suite")
	}

	dsn := os.Getenv("POSTGRES_TEST_URL")
	if dsn == "" {
		dsn = util.DefaultTestDsn
	}

	paymentCodeSuite := &inquiryTestSuite{
		util.Suite{
			DSN:                     dsn,
			MigrationLocationFolder: "migrations",
		},
	}

	suite.Run(t, paymentCodeSuite)
}

func (s inquiryTestSuite) BeforeTest(suiteName, testName string) {
	ok, err := s.Migration.Up()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s inquiryTestSuite) AfterTest(suiteName, testName string) {
	ok, err := s.Migration.Down()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func createMockInquiry() *domain.Inquiry {
	return &domain.Inquiry{
		ID:     uuid.UUID{},
		PaymentCode:   "test",
		TransactionID: "test_trxid",
	}
}

func (s inquiryTestSuite) TestCreateInquiry() {
	repo := repository.NewInquiryRepository(s.DBConn)

	testCases := []struct {
		desc        string
		ctx         context.Context
		reqBody     *domain.Inquiry
		expectedErr error
	}{
		{
			desc:        "success-create-inquiry",
			ctx:         context.Background(),
			reqBody:     createMockInquiry(),
			expectedErr: nil,
		},
		{
			desc: "error-create-payment-code-context-timeout",
			ctx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*-1))
				_ = cancel
				return ctx
			}(),
			reqBody:     createMockInquiry(),
			expectedErr: context.DeadlineExceeded,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			err := repo.Create(tC.ctx, tC.reqBody)
			if tC.expectedErr != nil {
				s.Require().Equal(tC.expectedErr, err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotEmpty(tC.reqBody.ID)
			inserted, err := repo.GetByTransactionID(tC.ctx, tC.reqBody.TransactionID)
			s.Require().NoError(err)
			reqBody := tC.reqBody

			s.Require().Equal(reqBody.TransactionID, inserted.TransactionID)
			s.Require().Equal(reqBody.PaymentCode, inserted.PaymentCode)
		})
	}
}

func (s inquiryTestSuite) TestGetByTransactionID() {
	repo := repository.NewInquiryRepository(s.DBConn)
	inquiry := createMockInquiry()

	testCases := []struct {
		desc           string
		reqBody        string
		expectedErr    error
		expectedResult *domain.Inquiry
	}{
		{
			desc:           "success-get-inquiry",
			reqBody:        inquiry.ID.String(),
			expectedErr:    nil,
			expectedResult: inquiry,
		},
		{
			desc:           "inquiry-doesnt-exist",
			reqBody:        "7072cd17-b1e7-4f21-98e2-33c81b3b17ad",
			expectedErr:    nil,
			expectedResult: nil,
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {

			_ = repo.Create(context.TODO(), inquiry)
			inserted, err := repo.GetByTransactionID(context.TODO(), tC.reqBody)
			s.Require().NoError(err)
			if inserted != (domain.Inquiry{}) {
				s.Require().Equal(inquiry.TransactionID, inserted.TransactionID)
				s.Require().Equal(inquiry.PaymentCode, inserted.PaymentCode)
			}
		})
	}

}
