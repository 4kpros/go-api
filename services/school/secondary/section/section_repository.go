package section

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/secondary/section/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(section *model.Section) (*model.Section, error) {
	result := *section
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(sectionID int64, userID int64, section *model.Section) (*model.Section, error) {
	tempSection, err := repository.GetById(sectionID, userID)
	if err != nil || tempSection == nil || tempSection.ID != sectionID {
		return nil, err
	}

	result := &model.Section{}
	return result, repository.Db.Model(result).Where("id = ?", sectionID).Updates(
		map[string]interface{}{
			"name":        section.Name,
			"description": section.Description,
		},
	).Error
}

func (repository *Repository) Delete(sectionID int64, userID int64) (int64, error) {
	tempSection, err := repository.GetById(sectionID, userID)
	if err != nil || tempSection == nil || tempSection.ID != sectionID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", sectionID).Delete(&model.Section{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(sectionID int64, userID int64) (*model.Section, error) {
	result := &model.Section{}
	return result, repository.Db.Model(&model.Section{}).
		Select("sections.*").
		Joins("left join school_directors on sections.school_id = school_directors.id").
		Where("sections.id = ?", sectionID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(section *model.Section) (*model.Section, error) {
	result := &model.Section{}
	return result, repository.Db.Where(section).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Section, error) {
	var result []model.Section
	return result, repository.Db.Model(&model.Section{}).
		Select("sections.*").
		Joins("left join school_directors on sections.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
