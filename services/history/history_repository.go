package history

import (
	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/history/model"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	Db                  *gorm.DB
	PermissionTableName string
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{Db: db, PermissionTableName: constants.PERMISSION_TABLE_NAME_HISTORY}
}

func (repository *HistoryRepository) Create(history *model.History) (*model.History, error) {
	result := *history
	return &result, repository.Db.Create(&result).Error
}

func (repository *HistoryRepository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.History, error) {
	result := []model.History{}
	return result, repository.Db.Scopes(utils.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
