package class

import (
	"fmt"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/class/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(class *model.Class) (*model.Class, error) {
	result := *class
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(classID int64, userID int64, class *model.Class) (*model.Class, error) {
	tempClass, err := repository.GetById(classID, userID)
	if err != nil || tempClass == nil || tempClass.ID != classID {
		return nil, err
	}

	result := &model.Class{}
	return result, repository.Db.Model(result).Where("id = ?", classID).Updates(
		map[string]interface{}{
			"name":        class.Name,
			"description": class.Description,
		},
	).Error
}

func (repository *Repository) Delete(classID int64, userID int64) (int64, error) {
	tempClass, err := repository.GetById(classID, userID)
	if err != nil || tempClass == nil || tempClass.ID != classID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", classID).Delete(&model.Class{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(classID int64, userID int64) (*model.Class, error) {
	result := &model.Class{}
	return result, repository.Db.Model(&model.Class{}).
		Select("classs.*").
		Joins("left join school_directors on classs.school_id = school_directors.id").
		Where("classs.id = ?", classID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(class *model.Class) (*model.Class, error) {
	result := &model.Class{}
	return result, repository.Db.Where(class).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Class, error) {
	var result []model.Class
	var condition string = ""
	if filter != nil && len(filter.Search) >= 1 {
		condition = fmt.Sprintf(
			"WHERE name ILIKE %s OR WHERE description ILIKE %s",
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"classes",
			condition,
			pagination,
			filter,
		),
	).Find(&result).Error

	// return result, repository.Db.Model(&model.Class{}).
	// 	Select("classs.*").
	// 	Joins("left join school_directors on classs.school_id = school_directors.id").
	// 	Where("school_directors.user_id = ?", userID).
	// 	Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
