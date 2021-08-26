// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	context "context"
	domain "frieda-golang-training-beginner/domain"

	mock "github.com/stretchr/testify/mock"
)

// IPaymentRepository is an autogenerated mock type for the IPaymentRepository type
type IPaymentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, payment
func (_m *IPaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	ret := _m.Called(ctx, payment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Payment) error); ok {
		r0 = rf(ctx, payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}