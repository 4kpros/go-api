package document

import (
	"fmt"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/document/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(document *model.Document) (*model.Document, error) {
	result := *document
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Delete(documentID int64) (int64, error) {
	result := repository.Db.Where("id = ?", documentID).Delete(&model.Document{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(documentID int64) (*model.Document, error) {
	result := &model.Document{}
	return result, repository.Db.Where("id = ?", documentID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(document *model.Document) (*model.Document, error) {
	result := &model.Document{}
	return result, repository.Db.Where(document).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.Document, error) {
	var result []model.Document
	var condition string = ""
	if filter != nil && len(filter.Search) >= 1 {
		condition = fmt.Sprintf(
			"WHERE name ILIKE %s OR WHERE description ILIKE %s OR WHERE type ILIKE %s",
			filter.Search,
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"documents",
			condition,
			pagination,
			filter,
		),
	).Find(&result).Error
}
