package repository_test

import (
	"context"
	"frieda-golang-training-beginner/domain"
	"frieda-golang-training-beginner/payment/repository"
	"frieda-golang-training-beginner/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type paymentTestSuite struct {
	util.Suite
}

func TestSuitePayment(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test Suite")
	}

	dsn := os.Getenv("POSTGRES_TEST_URL")
	if dsn == "" {
		dsn = util.DefaultTestDsn
	}

	paymentCodeSuite := &paymentTestSuite{
		util.Suite{
			DSN:                     dsn,
			MigrationLocationFolder: "migrations",
		},
	}

	suite.Run(t, paymentCodeSuite)
}

func (s paymentTestSuite) BeforeTest(suiteName, testName string) {
	ok, err := s.Migration.Up()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s paymentTestSuite) AfterTest(suiteName, testName string) {
	ok, err := s.Migration.Down()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func createMockPayment() *domain.Payment {
	return &domain.Payment{
		ID:     uuid.UUID{},
		PaymentCode:   "test",
		TransactionID: "test_trxid",
		Name: "test",
		Amount: 10000,
	}
}

func (s paymentTestSuite) TestCreatePayment() {
	repo := repository.NewPaymentRepository(s.DBConn)

	testCases := []struct {
		desc        string
		ctx         context.Context
		reqBody     *domain.Payment
		expectedErr error
	}{
		{
			desc:        "success-create-payment",
			ctx:         context.Background(),
			reqBody:     createMockPayment(),
			expectedErr: nil,
		},
		{
			desc: "error-create-payment-code-context-timeout",
			ctx: func() context.Context {
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*-1))
				_ = cancel
				return ctx
			}(),
			reqBody:     createMockPayment(),
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
		})
	}
}

