package history

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/history/model"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	Create(history *model.History) error
	GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.History, error)
}

type HistoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewHistoryRepositoryImpl(db *gorm.DB) HistoryRepository {
	return &HistoryRepositoryImpl{Db: db}
}

func (repository *HistoryRepositoryImpl) Create(history *model.History) error {
	return repository.Db.Create(history).Error
}

func (repository *HistoryRepositoryImpl) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.History, error) {
	var histories = []model.History{}
	var result = repository.Db.Scopes(utils.PaginationScope(histories, pagination, filter, repository.Db)).Find(histories)
	return histories, result.Error
}
