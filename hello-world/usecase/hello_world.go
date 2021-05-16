package usecase

import (
	"frieda-golang-training-beginner/domain"
)

type helloworldUsecase struct {}

func (h helloworldUsecase) Get() domain.HelloWorldResponsePayload {
	return domain.HelloWorldResponsePayload{Message: "Hello, world!"}
}

func NewHelloWorldUsecase() domain.HelloWorldUsecase {
	return &helloworldUsecase{}
}
