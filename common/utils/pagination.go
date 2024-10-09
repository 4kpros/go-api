package utils

import (
	"strings"

	"github.com/4kpros/go-api/common/types"
	"gorm.io/gorm"
)

func PaginationScope(model interface{}, pagination *types.Pagination, filters *types.Filter, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var count int64
	db.Model(model).Count(&count)
	pagination.UpdateFields(&count)
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.Offset).Limit(pagination.Limit).Order(filters.OrderBy + " " + filters.Sort)
	}
}

func GetPaginationFiltersFromQuery(filter *types.Filter, pagination *types.PaginationRequest) (*types.Pagination, *types.Filter) {
	page := pagination.Page
	limit := pagination.Limit

	search := filter.Search
	orderBy := filter.OrderBy
	sort := filter.Sort

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}
	if len(strings.TrimSpace(orderBy)) <= 0 {
		orderBy = "updated_at"
	}
	if sort != "asc" {
		sort = "desc"
	}

	return NewPaginationData(&page, &limit), NewFiltersData(&search, &orderBy, &sort)
}

func NewPaginationData(page *int, limit *int) *types.Pagination {
	var offset = 0
	if *page > 1 {
		offset = (*page - 1) * *limit
	}
	return &types.Pagination{
		CurrentPage:  *page,
		NextPage:     *page,
		PreviousPage: *page,
		TotalPages:   0,
		Count:        0,
		Limit:        *limit,
		Offset:       offset,
	}
}

func NewFiltersData(search *string, orderBy *string, sort *string) *types.Filter {
	return &types.Filter{
		Search:  *search,
		OrderBy: *orderBy,
		Sort:    *sort,
	}
}
