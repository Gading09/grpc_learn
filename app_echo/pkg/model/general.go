package model

type GetRequest struct {
	Query Query `json:"query"`
	Where Where `json:"where"`
}
type Query struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
	Field string `json:"field"`
}

type Where struct {
	Filter map[string]string `json:"filter"`
}
