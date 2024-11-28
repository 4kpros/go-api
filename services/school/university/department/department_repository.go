package department

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/department/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(department *model.Department) (*model.Department, error) {
	result := *department
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, userID int64, department *model.Department) (*model.Department, error) {
	tempDepartment, err := repository.GetById(id, userID)
	if err != nil || tempDepartment == nil || tempDepartment.ID != id {
		return nil, err
	}

	result := &model.Department{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name":        department.Name,
			"description": department.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64, userID int64) (int64, error) {
	tempDepartment, err := repository.GetById(id, userID)
	if err != nil || tempDepartment == nil || tempDepartment.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Department{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64, userID int64) (*model.Department, error) {
	result := &model.Department{}
	return result, repository.Db.Model(&model.Department{}).
		Select("departments.*").
		Joins("left join school_directors on departments.school_id = school_directors.id").
		Where("departments.id = ?", id).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(department *model.Department) (*model.Department, error) {
	result := &model.Department{}
	return result, repository.Db.Where(department).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Department, error) {
	var result []model.Department
	return result, repository.Db.Model(&model.Department{}).
		Select("departments.*").
		Joins("left join school_directors on departments.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
