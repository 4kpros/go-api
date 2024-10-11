package utils

import (
	"fmt"
	"strings"

	"github.com/4kpros/go-api/common/types"
	"gorm.io/gorm"
)

// Return *gorm.DB pointer with applied search, filter and pagination
//
// - Search performs a full-text search on every row of the table. It searches all fields.
// The official documentation for full text search in PostgreSQL can be found here
// https://www.postgresql.org/docs/17/textsearch-tables.html#TEXTSEARCH-TABLES-SEARCH
//
// - Filter applies an ORDER BY clause to the specified field name,
// sorting the results in ascending or descending based on the Sort parameter.
//
// - Pagination applies an offset and limit to the results, determining which subset of data to display.
func PaginationScope(model any, pagination *types.Pagination, filter *types.Filter, db *gorm.DB) func(*gorm.DB) *gorm.DB {
	var count int64
	db.Model(model).Count(&count)
	pagination.UpdateFields(&count)
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.Offset).Limit(pagination.Limit).Order(
			fmt.Sprintf("%s %s", filter.OrderBy, filter.Sort),
		).Where("to_tsvector(body) @@ to_tsquery(?)", filter.Search)
	}
}

// Check the entries and return the corrected ones.
func GetPaginationFiltersFromQuery(filter *types.Filter, pagination *types.PaginationRequest) (*types.Pagination, *types.Filter) {
	page := pagination.Page
	limit := pagination.Limit

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}
	if len(strings.TrimSpace(filter.OrderBy)) <= 0 {
		filter.OrderBy = "updated_at"
	}
	if filter.Sort != "asc" {
		filter.Sort = "desc"
	}

	return NewPaginationData(page, limit), filter
}

// The user passes a pagination request, specifying the desired page and limit.
// We validate the inputs and return a new pagination object with the applied settings:
// current page, next page, previous page, total pages, count, limit, and offset.
func NewPaginationData(page int, limit int) *types.Pagination {
	var offset = 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	return &types.Pagination{
		CurrentPage:  page,
		NextPage:     page,
		PreviousPage: page,
		TotalPages:   0,
		Count:        0,
		Limit:        limit,
		Offset:       offset,
	}
}
