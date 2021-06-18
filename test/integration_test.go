package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"frieda-golang-training-beginner/domain"
	http2 "frieda-golang-training-beginner/payment_code/directory/http"
	"frieda-golang-training-beginner/payment_code/repository"
	"frieda-golang-training-beginner/payment_code/usecase"
	"frieda-golang-training-beginner/util"
	"github.com/google/uuid"
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

func (s paymentCodeTestSuite) TestGet()  {
	IDstr := "f3c3f48e-41e6-40e1-afe3-95ab9529c769"
	var uuidReq [16]byte
	copy(uuidReq[:], IDstr)

	paymentCode := domain.PaymentCode{
		ID:          uuidReq,
		PaymentCode: "1",
		Name:        "Fr",
		Status:      "ACTIVE",
	}

	testCases := []struct {
		desc               string
		ID                        string
		statusCodeExpected        int
		returnPaymentCodeExpected domain.GetPaymentCodeResponsePayload
	}{
		{
			desc: "success-get-payment-codes",
			ID: IDstr,
			statusCodeExpected: http.StatusOK,
			returnPaymentCodeExpected: domain.GetPaymentCodeResponsePayload{
				ID:          uuidReq,
				PaymentCode: "1",
				Name:        "Fr",
				Status:      "ACTIVE",
			},
		},
		{
			desc:               "not-found",
			ID:                 uuid.New().String(),
			statusCodeExpected: http.StatusOK,
			returnPaymentCodeExpected: domain.GetPaymentCodeResponsePayload{},
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			repo := repository.NewPaymentCodeRepository(s.DBConn)
			service := usecase.NewPaymentCodeUsecase(repo, 5)
			repo.Create(context.TODO(), &paymentCode)
			e := echo.New()

			http2.NewPaymentCodeHandler(e, service)
			req := httptest.NewRequest("GET", fmt.Sprintf("/payment-codes/%s", tC.ID), nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			res, _ := json.Marshal(tC.returnPaymentCodeExpected)
			s.Require().Equal(tC.statusCodeExpected, rec.Code)
			s.Require().Equal(res, rec.Body.Bytes())

		})
	}
}
