package pupil

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/pupil/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(pupil *model.Pupil) (*model.Pupil, error) {
	result := *pupil
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(pupilID int64, userID int64, pupil *model.Pupil) (*model.Pupil, error) {
	tempPupil, err := repository.GetById(pupilID, userID)
	if err != nil || tempPupil == nil || tempPupil.ID != pupilID {
		return nil, err
	}

	result := &model.Pupil{}
	return result, repository.Db.Model(result).Where("school_id = ?", pupil.SchoolID).Updates(
		map[string]interface{}{
			"UserID":  pupil.UserID,
			"ClassID": pupil.ClassID,
		},
	).Error
}

func (repository *Repository) Delete(pupilID int64, userID int64) (int64, error) {
	tempPupil, err := repository.GetById(pupilID, userID)
	if err != nil || tempPupil == nil || tempPupil.ID != pupilID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", pupilID).Delete(&model.Pupil{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(pupilID int64, userID int64) (*model.Pupil, error) {
	result := &model.Pupil{}
	return result, repository.Db.Model(&model.Pupil{}).
		Select("pupils.*").
		Joins("left join school_directors on pupils.school_id = school_directors.id").
		Where("pupils.id = ?", pupilID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(pupil *model.Pupil) (*model.Pupil, error) {
	result := &model.Pupil{}
	return result, repository.Db.Where(pupil).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Pupil, error) {
	var result []model.Pupil
	return result, repository.Db.Model(&model.Pupil{}).
		Select("pupils.*").
		Joins("left join school_directors on pupils.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
