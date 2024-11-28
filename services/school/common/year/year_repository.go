package year

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/year/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(year *model.Year) (*model.Year, error) {
	result := *year
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, year *model.Year) (*model.Year, error) {
	result := &model.Year{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"start_date": year.StartDate,
			"end_date":   year.EndDate,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Year{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.Year, error) {
	result := &model.Year{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(year *model.Year) (*model.Year, error) {
	result := &model.Year{}
	return result, repository.Db.Where(year).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Year, error) {
	var result []model.Year
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
