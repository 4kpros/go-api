package school

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/admin/school/model"
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

func (repository *Repository) AddDirector(schoolId int64, userId int64) (*model.SchoolDirector, error) {
	result := &model.SchoolDirector{
		SchoolId: schoolId,
		UserId:   userId,
	}
	return result, repository.Db.Create(result).Error
}

func (repository *Repository) Update(id int64, school *model.School) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name": school.Name,
			"type": school.Type,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.School{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteDirector(schoolId int64, userId int64) (int64, error) {
	result := repository.Db.Where("schoolId = ?", schoolId).Where("userId = ?", userId).Delete(&model.SchoolDirector{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByName(name string) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *Repository) GetDirector(schoolId int64, userId int64) (*model.SchoolDirector, error) {
	result := &model.SchoolDirector{}
	return result, repository.Db.Where("schoolId = ?", schoolId).Where("userId = ?", userId).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.School, error) {
	var result []model.School
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
