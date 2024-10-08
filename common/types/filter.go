package types

type Filter struct {
	Search  string `json:"search" query:"search" doc:"Search keyword" example:""`
	OrderBy string `json:"orderBy" query:"orderBy" doc:"Filter by field name(case sensitive)" example:"id"`
	Sort    string `json:"sort" query:"sort" doc:"Sort asc or desc" example:"desc"`
}
