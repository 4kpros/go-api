package department

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/university/department/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityDepartment) (*model.UniversityDepartment, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.UniversityDepartment) (*model.UniversityDepartment, error) {
	result := &model.UniversityDepartment{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"school_id":   item.SchoolID,
			"name":        item.Name,
			"description": item.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityDepartment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.UniversityDepartment{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(id int64) (*model.UniversityDepartment, error) {
	result := &model.UniversityDepartment{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetBySchoolIDFacultyIDName(schoolID int64, facultyID int64, name string) (*model.UniversityDepartment, error) {
	result := &model.UniversityDepartment{}
	return result, repository.Db.Preload(clause.Associations).Where("school_id = ?", schoolID).Where("faculty_id = ?", facultyID).Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityDepartment, err error) {
	result = make([]model.UniversityDepartment, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE dps.school_id = '%d'", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(dps.id AS TEXT) = '%s' OR dps.name ILIKE '%s' OR dps.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR fcs.name ILIKE '%s' OR fcs.description ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)

		if strings.HasPrefix(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, tempWhere)
		} else {
			where = fmt.Sprintf("WHERE %s", tempWhere)
		}
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT dps.id, dps.name, dps.description, dps.school_id, dps.faculty_id"+
				", dps.created_at, dps.updated_at FROM university_departments dps "+
				"LEFT JOIN schools ON dps.school_id = schools.id "+
				"LEFT JOIN university_faculties AS fcs ON dps.faculty_id = fcs.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
