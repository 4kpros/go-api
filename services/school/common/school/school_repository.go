package school

import (
	"gorm.io/gorm"

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

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.School, error) {
	var result []model.School
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}

func (repository *Repository) GetAllByUserID(userID int64, filter *types.Filter, pagination *types.Pagination) ([]model.School, error) {
	var result []model.School
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Where("user_id = ?", userID).Find(result).Error
}

func (repository *Repository) GetAllDirectors(filter *types.Filter, pagination *types.Pagination) ([]model.SchoolDirector, error) {
	var result []model.SchoolDirector
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
