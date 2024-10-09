package types

type Filter struct {
	Search  string `json:"search" query:"search" required:"false" doc:"Search keyword" example:""`
	OrderBy string `json:"orderBy" query:"orderBy" required:"false" doc:"Filter by field name(case sensitive)" example:"id"`
	Sort    string `json:"sort" query:"sort" required:"false" doc:"Sort asc or desc" example:"desc"`
}
