package usecase

import (
	"frieda-golang-training-beginner/domain"
)

type HealthUsecase struct{}

func (h HealthUsecase) Get() domain.HealthResponsePayload {
	return domain.HealthResponsePayload{Status: "healthy"}
}


