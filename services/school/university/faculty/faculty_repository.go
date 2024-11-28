package faculty

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/faculty/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(faculty *model.Faculty) (*model.Faculty, error) {
	result := *faculty
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, faculty *model.Faculty) (*model.Faculty, error) {
	result := &model.Faculty{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name":        faculty.Name,
			"description": faculty.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Faculty{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.Faculty, error) {
	result := &model.Faculty{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(faculty *model.Faculty) (*model.Faculty, error) {
	result := &model.Faculty{}
	return result, repository.Db.Where(faculty).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Faculty, error) {
	var result []model.Faculty
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
