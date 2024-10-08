package types

type ErrorResponse struct {
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Filter     *Filter     `json:"filter"`
	Pagination *Pagination `json:"pagination"`
}
