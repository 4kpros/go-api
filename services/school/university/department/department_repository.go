package department

import (
	"fmt"

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

func (repository *Repository) Update(departmentID int64, userID int64, department *model.Department) (*model.Department, error) {
	tempDepartment, err := repository.GetById(departmentID, userID)
	if err != nil || tempDepartment == nil || tempDepartment.ID != departmentID {
		return nil, err
	}

	result := &model.Department{}
	return result, repository.Db.Model(result).Where("id = ?", departmentID).Updates(
		map[string]interface{}{
			"name":        department.Name,
			"description": department.Description,
		},
	).Error
}

func (repository *Repository) Delete(departmentID int64, userID int64) (int64, error) {
	tempDepartment, err := repository.GetById(departmentID, userID)
	if err != nil || tempDepartment == nil || tempDepartment.ID != departmentID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", departmentID).Delete(&model.Department{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(departmentID int64, userID int64) (*model.Department, error) {
	result := &model.Department{}
	return result, repository.Db.Model(&model.Department{}).
		Select("departments.*").
		Joins("left join school_directors on departments.school_id = school_directors.id").
		Where("departments.id = ?", departmentID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(department *model.Department) (*model.Department, error) {
	result := &model.Department{}
	return result, repository.Db.Where(department).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Department, error) {
	var result []model.Department
	var selection string = fmt.Sprintf(
		"departments.* from departments left join school_directors on departments.school_id = school_directors.id WHERE school_directors.user_id = %d",
		userID,
	)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE name ILIKE %s OR WHERE description ILIKE %s",
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			selection,
			where,
			pagination,
			filter,
		),
	).Find(&result).Error
}
