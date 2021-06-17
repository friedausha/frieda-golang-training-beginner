package http_test

import (
	"frieda-golang-training-beginner/domain"
	mocks "frieda-golang-training-beginner/mock"
	http2 "frieda-golang-training-beginner/payment_code/directory/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreate(t *testing.T)  {
	e := echo.New()
	testCases := []struct {
		desc               string
		svc                *mocks.IPaymentCodeUsecase
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
			svc: func() *mocks.IPaymentCodeUsecase {
				mockSvc := new(mocks.IPaymentCodeUsecase)
				mockSvc.On("Create", mock.Anything, mock.Anything).Once().
					Return(domain.PaymentCode{Name: "Fr"}, nil)
				return mockSvc
			}(),
		},
		{
			desc: "error-create-payment-statement-bad-request-body",
			reqBody: strings.NewReader(
				`{}`),
			statusCodeExpected: http.StatusBadRequest,
			svc: func() *mocks.IPaymentCodeUsecase {
				mockSvc := new(mocks.IPaymentCodeUsecase)
				return mockSvc
			}(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			http2.NewPaymentCodeHandler(e, tC.svc)
			req := httptest.NewRequest("POST", "/payment-codes", tC.reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			require.Equal(t, tC.statusCodeExpected, rec.Code)
			tC.svc.AssertExpectations(t)
		})
	}

}
