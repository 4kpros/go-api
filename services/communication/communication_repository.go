package communication

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/communication/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(communication *model.Communication) (*model.Communication, error) {
	result := *communication
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Delete(roleID int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", roleID).Delete(&model.Communication{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Communication{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(communicationID int64) (*model.Communication, error) {
	result := &model.Communication{}
	return result, repository.Db.Where("id = ?", communicationID).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.Communication, err error) {
	result = make([]model.Communication, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE subject ILIKE '%s' OR message ILIKE '%s'",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM communications",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
