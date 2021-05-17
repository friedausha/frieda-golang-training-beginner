package usecase

import (
	"frieda-golang-training-beginner/domain"
)

type healthUsecase struct{}

func (h healthUsecase) Get() domain.HealthResponsePayload {
	return domain.HealthResponsePayload{Status: "healthy"}
}
