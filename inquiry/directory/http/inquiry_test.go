package http_test

import (
	"frieda-golang-training-beginner/domain"
	http2 "frieda-golang-training-beginner/inquiry/directory/http"
	mocks "frieda-golang-training-beginner/mock"
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
		svc                *mocks.IInquiryUsecase
		reqBody            io.Reader
		statusCodeExpected int
	}{
		{
			desc: "success-create-inquiry",
			reqBody: strings.NewReader(
				`{
					"payment_code": "714912",
					"transaction_id": "1"
					}
				`),
			statusCodeExpected: http.StatusCreated,
			svc: func() *mocks.IInquiryUsecase {
				mockSvc := new(mocks.IInquiryUsecase)
				mockSvc.On("Create", mock.Anything, mock.Anything).Once().
					Return(domain.Inquiry{TransactionID: "Fr"}, nil)
				return mockSvc
			}(),
		},
		{
			desc: "error-create-inquiry-bad-request-body",
			reqBody: strings.NewReader(
				`{}`),
			statusCodeExpected: http.StatusBadRequest,
			svc: func() *mocks.IInquiryUsecase {
				mockSvc := new(mocks.IInquiryUsecase)
				return mockSvc
			}(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			http2.NewPaymentCodeHandler(e, tC.svc)
			req := httptest.NewRequest("POST", "/inquiry", tC.reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			require.Equal(t, tC.statusCodeExpected, rec.Code)
		})
	}
}
