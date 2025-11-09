package dto

type SuccessResponse struct {
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data,omitempty"`
}
