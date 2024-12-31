package class

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/highschool/class/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.HighschoolClass) (*model.HighschoolClass, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.HighschoolClass) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	fmt.Println(item)
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"school_id":    item.SchoolID,
			"specialty_id": item.SpecialtyID,
			"name":         item.Name,
			"description":  item.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolClass{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.HighschoolClass{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(id int64) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetBySchoolIDSpecialtyIDName(schoolID int64, specialtyID int64, name string) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	return result, repository.Db.Preload(clause.Associations).Where("school_id = ?", schoolID).Where("specialty_id = ?", specialtyID).Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolClass, err error) {
	result = make([]model.HighschoolClass, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE classes.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(classes.id AS TEXT) = '%s' OR classes.name ILIKE '%s' OR classes.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR specialties.name ILIKE '%s' OR specialties.description ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)

		if strings.HasPrefix(where, "WHERE") {
			where = fmt.Sprintf("%s AND (%s)", where, tempWhere)
		} else {
			where = fmt.Sprintf("WHERE %s", tempWhere)
		}
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT classes.id, classes.name, classes.description, classes.school_id, classes.specialty_id"+
				", classes.created_at, classes.updated_at FROM highschool_classes classes "+
				"LEFT JOIN schools ON classes.school_id = schools.id "+
				"LEFT JOIN highschool_specialties AS specialties ON classes.specialty_id = specialties.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
