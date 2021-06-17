package integration_test

import (
	http2 "frieda-golang-training-beginner/payment_code/directory/http"
	"frieda-golang-training-beginner/payment_code/repository"
	"frieda-golang-training-beginner/payment_code/usecase"
	"frieda-golang-training-beginner/util"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type paymentCodeTestSuite struct {
	util.Suite
}

func TestSuitePaymentCode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test Suite")
	}

	dsn := os.Getenv("POSTGRES_TEST_URL")
	if dsn == "" {
		dsn = util.DefaultTestDsn
	}

	paymentCodeSuite := &paymentCodeTestSuite{
		util.Suite{
			DSN:                     dsn,
			MigrationLocationFolder: "../payment_code/repository/migrations",
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

func (s paymentCodeTestSuite) TestCreate() {
	testCases := []struct {
		desc               string
		reqBody            io.Reader
		statusCodeExpected int
	}{
		{
			desc: "success-create-payment-codes",
			reqBody: strings.NewReader(
				`{
					"payment_code": "714912",
					"status": "ACTIVE",
					"name": "Fr"
					}
				`),
			statusCodeExpected: http.StatusCreated,
		},
		{
			desc: "error-create-payment-statement-bad-request-body",
			reqBody: strings.NewReader(
				`{}`),
			statusCodeExpected: http.StatusBadRequest,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			repo := repository.NewPaymentCodeRepository(s.DBConn)
			service := usecase.NewPaymentCodeUsecase(repo, 5)

			e := echo.New()

			http2.NewPaymentCodeHandler(e, service)
			req := httptest.NewRequest("POST", "/payment-codes", tC.reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			s.Require().Equal(tC.statusCodeExpected, rec.Code)

		})
	}

}
