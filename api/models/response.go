package models

// SuccessResponse is used to respond to successful requests.
type SuccessResponse struct {
	Code    int
	Ok      bool
	Payload interface{}
}

// ErrorResponse is used to respond to failed requests.
type ErrorResponse struct {
	Code    int
	Ok      bool
	Message string
}
