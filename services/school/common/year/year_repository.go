package year

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
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
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(yearID int64, year *model.Year) (*model.Year, error) {
	result := &model.Year{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", yearID).Updates(
		map[string]interface{}{
			"name":       year.Name,
			"start_date": year.StartDate,
			"end_date":   year.EndDate,
			"school_id":  year.SchoolID,
		},
	).Error
}

func (repository *Repository) Delete(yearID int64) (int64, error) {
	result := repository.Db.Where("id = ?", yearID).Delete(&model.Year{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Year{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(yearID int64) (*model.Year, error) {
	result := &model.Year{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", yearID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByNameSchoolID(name string, schoolID int64) (*model.Year, error) {
	result := &model.Year{}
	return result, repository.Db.Preload(clause.Associations).Where("name = ?", name).Where("school_id = ?", schoolID).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.Year, err error) {
	result = make([]model.Year, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE CAST(id AS TEXT) = '%s' OR name ILIKE '%s' OR start_date ILIKE '%s' OR end_date ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM years",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
