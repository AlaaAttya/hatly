package createsession

type Request struct {
	Code     string `json:"code" validate:"required"`
	Username string `json:"username" validate:"required"`
}
