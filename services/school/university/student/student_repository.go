package student

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/student/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(student *model.Student) (*model.Student, error) {
	result := *student
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(studentID int64, userID int64, student *model.Student) (*model.Student, error) {
	tempStudent, err := repository.GetById(studentID, userID)
	if err != nil || tempStudent == nil || tempStudent.ID != studentID {
		return nil, err
	}

	result := &model.Student{}
	return result, repository.Db.Model(result).Where("school_id = ?", student.SchoolID).Updates(
		map[string]interface{}{
			"UserID":  student.UserID,
			"LevelID": student.LevelID,
		},
	).Error
}

func (repository *Repository) Delete(studentID int64, userID int64) (int64, error) {
	tempStudent, err := repository.GetById(studentID, userID)
	if err != nil || tempStudent == nil || tempStudent.ID != studentID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", studentID).Delete(&model.Student{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(studentID int64, userID int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Model(&model.Student{}).
		Select("students.*").
		Joins("left join school_directors on students.school_id = school_directors.id").
		Where("students.id = ?", studentID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(student *model.Student) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Where(student).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Student, error) {
	var result []model.Student
	return result, repository.Db.Model(&model.Student{}).
		Select("students.*").
		Joins("left join school_directors on students.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
