package domain

type HelloWorldResponsePayload struct {
	Message string `json:"message"`
}
type HelloWorldUsecase interface {
	Get() HelloWorldResponsePayload
}
