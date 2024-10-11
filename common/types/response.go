package types

type ErrorResponse struct {
	Message string `json:"message" required:"false"`
}

type PaginatedResponse struct {
	Filter     *Filter     `json:"filter" required:"false"`
	Pagination *Pagination `json:"pagination" required:"false"`
}

type DeletedResponse struct {
	AffectedRows int64 `json:"affectedRows" required:"false" doc:"Number of row affected with this delete" example:"1"`
}
