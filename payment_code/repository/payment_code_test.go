package repository_test

import (
	"context"
	"frieda-golang-training-beginner/domain"
	"frieda-golang-training-beginner/payment_code/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type paymentCodeTestSuite struct {
	repository.Suite
}

func TestSuitePaymentCode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test Suite")
	}

	dsn := os.Getenv("POSTGRES_TEST_URL")
	if dsn == "" {
		dsn = repository.DefaultTestDsn
	}

	paymentCodeSuite := &paymentCodeTestSuite{
		repository.Suite{
			DSN:                     dsn,
			MigrationLocationFolder: "migrations",
		},
	}

	suite.Run(t, paymentCodeSuite)
}

func (s paymentCodeTestSuite) BeforeTest(suiteName, testName string) {
	ok, err := s.Migration.Up()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s paymentCodeTestSuite) AfterTest(suiteName, testName string) {
	ok, err := s.Migration.Down()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func createMockPaymentCodes() *domain.PaymentCode {
	return &domain.PaymentCode{
		ID:     uuid.UUID{},
		Name:   "Test",
		Status: "ACTIVE",
	}
}

func (s paymentCodeTestSuite) TestCreatePaymentCodes() {
	repo := repository.NewPaymentCodeRepository(s.DBConn)

	testCases := []struct {
		desc        string
		ctx         context.Context
		reqBody     *domain.PaymentCode
		expectedErr error
	}{
		{
			desc:        "success-create-payment-code",
			ctx:         context.Background(),
			reqBody:     createMockPaymentCodes(),
			expectedErr: nil,
		},
		{
			desc: "error-create-payment-code-context-timeout",
			ctx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*-1))
				_ = cancel
				return ctx
			}(),
			reqBody:     createMockPaymentCodes(),
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
			inserted, err := repo.GetByID(tC.ctx, tC.reqBody.ID.String())
			s.Require().NoError(err)
			reqBody := tC.reqBody

			s.Require().Equal(reqBody.Name, inserted.Name)
			s.Require().Equal(reqBody.Status, inserted.Status)
		})
	}
}

func (s paymentCodeTestSuite) TestGetByID() {
	repo := repository.NewPaymentCodeRepository(s.DBConn)
	paymentCode := createMockPaymentCodes()

	testCases := []struct {
		desc           string
		reqBody        string
		expectedErr    error
		expectedResult *domain.PaymentCode
	}{
		{
			desc:           "success-get-payment-code-by=id",
			reqBody:        paymentCode.ID.String(),
			expectedErr:    nil,
			expectedResult: paymentCode,
		},
		{
			desc:           "error-payment-code-doesnt-exist",
			reqBody:        "7072cd17-b1e7-4f21-98e2-33c81b3b17ad",
			expectedErr:    nil,
			expectedResult: nil,
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {

			_ = repo.Create(context.TODO(), paymentCode)
			inserted, err := repo.GetByID(context.TODO(), tC.reqBody)
			s.Require().NoError(err)
			if inserted != (domain.PaymentCode{}) {
				s.Require().Equal(paymentCode.Name, inserted.Name)
				s.Require().Equal(paymentCode.Status, inserted.Status)
			}
		})
	}

}
