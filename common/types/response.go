package types

type ErrorResponse struct {
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Filter     *Filter     `json:"filter"`
	Pagination *Pagination `json:"pagination"`
}

type DeleteResponse struct {
	AffectedRows int64 `json:"affectedRows" doc:"Number of row affected with this delete" example:"1"`
}
