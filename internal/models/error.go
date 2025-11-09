package models

// Error представляет структуру ошибки для API ответов
type Error struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
