package usecase

import (
	"frieda-golang-training-beginner/domain"
)

type HelloWorldUsecase struct {}

func (h HelloWorldUsecase) Get() domain.HelloWorldResponsePayload {
	return domain.HelloWorldResponsePayload{Message: "hello world"}
}