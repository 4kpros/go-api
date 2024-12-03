package domain

import (
	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/domain/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(domain *model.Domain) (*model.Domain, error) {
	result := *domain
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, userID int64, domain *model.Domain) (*model.Domain, error) {
	tempDomain, err := repository.GetById(id, userID)
	if err != nil || tempDomain == nil || tempDomain.ID != id {
		return nil, err
	}

	result := &model.Domain{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name":        domain.Name,
			"description": domain.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64, userID int64) (int64, error) {
	tempDomain, err := repository.GetById(id, userID)
	if err != nil || tempDomain == nil || tempDomain.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Domain{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64, userID int64) (*model.Domain, error) {
	result := &model.Domain{}
	return result, repository.Db.Model(&model.Domain{}).
		Select("domains.*").
		Joins("left join school_directors on domains.school_id = school_directors.id").
		Where("domains.id = ?", id).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(domain *model.Domain) (*model.Domain, error) {
	result := &model.Domain{}
	return result, repository.Db.Where(domain).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.Domain, error) {
	result := make([]model.Domain, 0)
	return result, repository.Db.Model(&model.Domain{}).
		Select("domains.*").
		Joins("left join school_directors on domains.school_id = school_directors.id").
		Where("school_directors.user_id = ?", userID).
		Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
