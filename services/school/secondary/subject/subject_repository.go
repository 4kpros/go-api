package subject

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/subject/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(subject *model.Subject) (*model.Subject, error) {
	result := *subject
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) AddProfessor(professor *model.SubjectProfessor) (*model.SubjectProfessor, error) {
	result := *professor
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(subjectID int64, userID int64, subject *model.Subject) (*model.Subject, error) {
	tempSubject, err := repository.GetById(subjectID, userID)
	if err != nil || tempSubject == nil || tempSubject.ID != subjectID {
		return nil, err
	}

	result := &model.Subject{}
	return result, repository.Db.Model(result).Where("id = ?", subjectID).Updates(
		map[string]interface{}{
			"name":         subject.Name,
			"description":  subject.Description,
			"coefficient":  subject.Coefficient,
			"program":      subject.Program,
			"requirements": subject.Requirements,
		},
	).Error
}

func (repository *Repository) Delete(subjectID int64, userID int64) (int64, error) {
	tempSubject, err := repository.GetById(subjectID, userID)
	if err != nil || tempSubject == nil || tempSubject.ID != subjectID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", subjectID).Delete(&model.Subject{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteProfessor(subjectID int64, userID int64) (int64, error) {
	professor, err := repository.GetProfessorById(subjectID, userID)
	if err != nil || professor == nil || professor.ID != subjectID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", subjectID).Where("user_id = ?", userID).Delete(&model.SubjectProfessor{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(subjectID int64, userID int64) (*model.Subject, error) {
	result := &model.Subject{}
	return result, repository.Db.Model(&model.Subject{}).
		Select("subjects.*").
		Joins("left join school_directors on subjects.school_id = school_directors.id").
		Where("subjects.id = ?", subjectID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetProfessorById(subjectProfessorID int64, userID int64) (*model.SubjectProfessor, error) {
	result := &model.SubjectProfessor{}
	return result, repository.Db.Where("id = ?", subjectProfessorID).Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(subject *model.Subject) (*model.Subject, error) {
	result := &model.Subject{}
	return result, repository.Db.Where(subject).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Subject, error) {
	var result []model.Subject
	return result, repository.Db.Model(&model.Subject{}).
		Select("subjects.*").
		Joins("left join school_directors on subjects.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
