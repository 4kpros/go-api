package history

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/history/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(history *model.History) (*model.History, error) {
	result := *history
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.History, error) {
	result := make([]model.History, 0)
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
