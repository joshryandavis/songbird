package stmodels

type ErrorResponse struct {
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
	Success bool `json:"success"`
}
