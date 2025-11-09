package dto

type ErrorResponse struct {
	Message string `json:"message" example:"status bad request"`
	Details string `json:"details" example:"status bad request"`
}
