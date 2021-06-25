package http_test

import (
	"fmt"
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

func TestCreate(t *testing.T) {
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
		})
	}
}

func TestGetPaymentCode(t *testing.T) {
	e := echo.New()
	IDstr := "f3c3f48e-41e6-40e1-afe3-95ab9529c769"
	var uuid [16]byte
	copy(uuid[:], IDstr)

	testCases := []struct {
		desc                      string
		svc                       *mocks.IPaymentCodeUsecase
		ID                        string
		statusCodeExpected        int
		returnPaymentCodeExpected string
	}{
		{
			desc:               "success-get-payment-codes",
			ID:                 IDstr,
			statusCodeExpected: http.StatusOK,
			svc: func() *mocks.IPaymentCodeUsecase {
				mockSvc := new(mocks.IPaymentCodeUsecase)
				mockSvc.On("Get", mock.Anything, mock.Anything).Once().
					Return(domain.GetPaymentCodeResponsePayload{Name: "Fr", ID: uuid, PaymentCode: "1", Status: "ACTIVE"}, nil)
				return mockSvc
			}(),
			returnPaymentCodeExpected: "{\"status\":\"ACTIVE\",\"id\":\"66336333-6634-3865-2d34-3165362d3430\",\"payment_code\":\"1\",\"name\":\"Fr\"}\n",
		},
		{
			desc:               "not-found",
			ID:                 IDstr,
			statusCodeExpected: http.StatusOK,
			svc: func() *mocks.IPaymentCodeUsecase {
				mockSvc := new(mocks.IPaymentCodeUsecase)
				mockSvc.On("Get", mock.Anything, mock.Anything).Once().
					Return(domain.GetPaymentCodeResponsePayload{}, nil)
				return mockSvc
			}(),
			returnPaymentCodeExpected: "{\"status\":\"\",\"id\":\"00000000-0000-0000-0000-000000000000\",\"payment_code\":\"\",\"name\":\"\"}\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			http2.NewPaymentCodeHandler(e, tC.svc)
			req := httptest.NewRequest("GET", fmt.Sprintf("/payment-codes/%s", tC.ID), nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			require.Equal(t, tC.statusCodeExpected, rec.Code)
			require.Equal(t, tC.returnPaymentCodeExpected, rec.Body.String())
		})
	}

}
