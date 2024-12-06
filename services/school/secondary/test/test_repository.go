package test

import (
	"fmt"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/test/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(test *model.Test) (*model.Test, error) {
	result := *test
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(testID int64, userID int64, test *model.Test) (*model.Test, error) {
	tempTest, err := repository.GetById(testID, userID)
	if err != nil || tempTest == nil || tempTest.ID != testID {
		return nil, err
	}

	result := &model.Test{}
	return result, repository.Db.Model(result).Where("id = ?", testID).Updates(
		map[string]interface{}{
			"type":        test.Type,
			"percentage":  test.Percentage,
			"description": test.Description,
		},
	).Error
}

func (repository *Repository) Delete(testID int64, userID int64) (int64, error) {
	tempTest, err := repository.GetById(testID, userID)
	if err != nil || tempTest == nil || tempTest.ID != testID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", testID).Delete(&model.Test{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(testID int64, userID int64) (*model.Test, error) {
	result := &model.Test{}
	return result, repository.Db.Model(&model.Test{}).
		Select("tests.*").
		Joins("left join school_directors on tests.school_id = school_directors.id").
		Where("tests.id = ?", testID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(test *model.Test) (*model.Test, error) {
	result := &model.Test{}
	return result, repository.Db.Where(test).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Test, error) {
	var result []model.Test
	var condition string = ""
	if filter != nil && len(filter.Search) >= 1 {
		condition = fmt.Sprintf(
			"WHERE type ILIKE %s OR WHERE description ILIKE %s",
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"tests",
			condition,
			pagination,
			filter,
		),
	).Find(&result).Error

	// return result, repository.Db.Model(&model.Test{}).
	// 	Select("tests.*").
	// 	Joins("left join school_directors on tests.school_id = school_directors.id").
	// 	Where("school_directors.user_id = ?", userID).
	// 	Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
