package domain

type HealthResponsePayload struct {
	Status        string     `json:"status"`
}
type HealthUsecase interface {
	Get() HealthResponsePayload
}
