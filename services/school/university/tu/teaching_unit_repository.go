package tu

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/tu/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(teachingUnit *model.TeachingUnit) (*model.TeachingUnit, error) {
	result := *teachingUnit
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) AddProfessor(professor *model.TeachingUnitProfessor) (*model.TeachingUnitProfessor, error) {
	result := *professor
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, userID int64, teachingUnit *model.TeachingUnit) (*model.TeachingUnit, error) {
	tempTeachingUnit, err := repository.GetById(id, userID)
	if err != nil || tempTeachingUnit == nil || tempTeachingUnit.ID != id {
		return nil, err
	}

	result := &model.TeachingUnit{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name":         teachingUnit.Name,
			"description":  teachingUnit.Description,
			"credit":       teachingUnit.Credit,
			"program":      teachingUnit.Program,
			"requirements": teachingUnit.Requirements,
		},
	).Error
}

func (repository *Repository) Delete(id int64, userID int64) (int64, error) {
	tempTeachingUnit, err := repository.GetById(id, userID)
	if err != nil || tempTeachingUnit == nil || tempTeachingUnit.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.TeachingUnit{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteProfessor(id int64, userID int64) (int64, error) {
	professor, err := repository.GetProfessorById(id, userID)
	if err != nil || professor == nil || professor.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Where("user_id = ?", userID).Delete(&model.TeachingUnitProfessor{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64, userID int64) (*model.TeachingUnit, error) {
	result := &model.TeachingUnit{}
	return result, repository.Db.Model(&model.TeachingUnit{}).
		Select("teaching_units.*").
		Joins("left join school_directors on teaching_units.school_id = school_directors.id").
		Where("teaching_units.id = ?", id).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetProfessorById(id int64, userID int64) (*model.TeachingUnitProfessor, error) {
	result := &model.TeachingUnitProfessor{}
	return result, repository.Db.Where("id = ?", id).Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(teachingUnit *model.TeachingUnit) (*model.TeachingUnit, error) {
	result := &model.TeachingUnit{}
	return result, repository.Db.Where(teachingUnit).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.TeachingUnit, error) {
	var result []model.TeachingUnit
	return result, repository.Db.Model(&model.TeachingUnit{}).
		Select("teaching_units.*").
		Joins("left join school_directors on teaching_units.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
