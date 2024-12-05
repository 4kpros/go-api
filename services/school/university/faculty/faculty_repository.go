package faculty

import (
	"fmt"

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

func (repository *Repository) Update(id int64, userID int64, faculty *model.Faculty) (*model.Faculty, error) {
	tempFaculty, err := repository.GetById(id, userID)
	if err != nil || tempFaculty == nil || tempFaculty.ID != id {
		return nil, err
	}

	result := &model.Faculty{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name":        faculty.Name,
			"description": faculty.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64, userID int64) (int64, error) {
	tempFaculty, err := repository.GetById(id, userID)
	if err != nil || tempFaculty == nil || tempFaculty.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Faculty{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64, userID int64) (*model.Faculty, error) {
	result := &model.Faculty{}
	return result, repository.Db.Model(&model.Faculty{}).
		Select("faculties.*").
		Joins("left join school_directors on faculties.school_id = school_directors.id").
		Where("faculties.id = ?", id).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(faculty *model.Faculty) (*model.Faculty, error) {
	result := &model.Faculty{}
	return result, repository.Db.Where(faculty).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Faculty, error) {
	var result []model.Faculty
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
			"faculties",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	// return result, repository.Db.Model(&model.Faculty{}).
	//
	//	Select("faculties.*").
	//	Joins("left join school_directors on faculties.school_id = school_directors.id").
	//	Where("school_directors.user_id = ?", userID).
	//	Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
