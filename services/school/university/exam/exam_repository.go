package exam

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/exam/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(exam *model.Exam) (*model.Exam, error) {
	result := *exam
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, userID int64, exam *model.Exam) (*model.Exam, error) {
	tempExam, err := repository.GetById(id, userID)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return nil, err
	}

	result := &model.Exam{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"type":        exam.Type,
			"percentage":  exam.Percentage,
			"description": exam.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64, userID int64) (int64, error) {
	tempExam, err := repository.GetById(id, userID)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Exam{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64, userID int64) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Model(&model.Exam{}).
		Select("exams.*").
		Joins("left join school_directors on exams.school_id = school_directors.id").
		Where("exams.id = ?", id).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(exam *model.Exam) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Where(exam).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Exam, error) {
	var result []model.Exam
	return result, repository.Db.Model(&model.Exam{}).
		Select("exams.*").
		Joins("left join school_directors on exams.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
