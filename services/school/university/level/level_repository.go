package level

import (
	"fmt"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/level/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(level *model.Level) (*model.Level, error) {
	result := *level
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(levelID int64, userID int64, level *model.Level) (*model.Level, error) {
	tempLevel, err := repository.GetById(levelID, userID)
	if err != nil || tempLevel == nil || tempLevel.ID != levelID {
		return nil, err
	}

	result := &model.Level{}
	return result, repository.Db.Model(result).Where("id = ?", levelID).Updates(
		map[string]interface{}{
			"name":        level.Name,
			"description": level.Description,
		},
	).Error
}

func (repository *Repository) Delete(levelID int64, userID int64) (int64, error) {
	tempLevel, err := repository.GetById(levelID, userID)
	if err != nil || tempLevel == nil || tempLevel.ID != levelID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", levelID).Delete(&model.Level{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(levelID int64, userID int64) (*model.Level, error) {
	result := &model.Level{}
	return result, repository.Db.Model(&model.Level{}).
		Select("levels.*").
		Joins("left join school_directors on levels.school_id = school_directors.id").
		Where("levels.id = ?", levelID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(level *model.Level) (*model.Level, error) {
	result := &model.Level{}
	return result, repository.Db.Where(level).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Level, error) {
	var result []model.Level
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
			"levels",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	// return result, repository.Db.Model(&model.Level{}).
	// 	Select("levels.*").
	// 	Joins("left join school_directors on levels.school_id = school_directors.id").
	// 	Where("school_directors.user_id = ?", userID).
	// 	Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
