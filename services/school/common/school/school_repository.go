package school

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/school/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(school *model.School) (*model.School, error) {
	result := *school
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) AddDirector(schoolID int64, userID int64) (*model.SchoolDirector, error) {
	result := &model.SchoolDirector{
		SchoolID: schoolID,
		UserID:   userID,
	}
	return result, repository.Db.Create(result).Error
}

func (repository *Repository) Update(schoolID int64, school *model.School) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Model(result).Where("id = ?", schoolID).Updates(
		map[string]interface{}{
			"name": school.Name,
			"type": school.Type,
		},
	).Error
}

func (repository *Repository) Delete(schoolID int64) (int64, error) {
	result := repository.Db.Where("id = ?", schoolID).Delete(&model.School{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteDirector(schoolID int64, userID int64) (int64, error) {
	result := repository.Db.Where("school_id = ?", schoolID).Where("user_id = ?", userID).Delete(&model.SchoolDirector{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(schoolID int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Where("id = ?", schoolID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByName(name string) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *Repository) GetDirector(schoolID int64, userID int64) (*model.SchoolDirector, error) {
	result := &model.SchoolDirector{}
	return result, repository.Db.Where("school_id = ?", schoolID).Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.School, err error) {
	result = make([]model.School, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE name ILIKE '%s' OR type ILIKE '%s'",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM schools",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllByUserID(userID int64, filter *types.Filter, pagination *types.Pagination) ([]model.School, error) {
	var result []model.School
	var condition string = ""
	if filter != nil && len(filter.Search) >= 1 {
		condition = fmt.Sprintf(
			"WHERE name ILIKE %s OR WHERE type ILIKE %s AND user_id = %d",
			filter.Search,
			filter.Search,
			userID,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"schools",
			condition,
			pagination,
			filter,
		),
	).Find(&result).Error
}

func (repository *Repository) GetAllDirectors(filter *types.Filter, pagination *types.Pagination) ([]model.SchoolDirector, error) {
	var result []model.SchoolDirector
	var condition string = ""
	if filter != nil && len(filter.Search) >= 1 {
		condition = ""
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"school_directors",
			condition,
			pagination,
			filter,
		),
	).Find(&result).Error
}
