package school

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/school/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.School) (*model.School, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateInfo(item *model.SchoolInfo) (*model.SchoolInfo, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateConfig(item *model.SchoolConfig) (*model.SchoolConfig, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.School) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"name": item.Name,
			"type": item.Type,
		},
	).Error
}

func (repository *Repository) UpdateConfigInfoIDs(id int64, configID int64, infoID int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"config_id": configID,
			"info_id":   infoID,
		},
	).Error
}

func (repository *Repository) UpdateInfo(id int64, item *model.SchoolInfo) (*model.SchoolInfo, error) {
	result := &model.SchoolInfo{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"full_name":   item.FullName,
			"description": item.Description,
			"slogan":      item.Slogan,

			"phone_number1": item.PhoneNumber1,
			"phone_number2": item.PhoneNumber2,
			"phone_number3": item.PhoneNumber3,

			"email1": item.Email1,
			"email2": item.Email2,
			"email3": item.Email3,

			"founder":    item.Founder,
			"founded_at": item.FoundedAt,

			"address":            item.Address,
			"location_longitude": item.LocationLongitude,
			"location_latitude":  item.LocationLatitude,

			"logo": item.Logo,

			"image1": item.Image1,
			"image2": item.Image2,
			"image3": item.Image3,
			"image4": item.Image4,
		},
	).Error
}

func (repository *Repository) UpdateConfig(id int64, item *model.SchoolConfig) (*model.SchoolConfig, error) {
	result := &model.SchoolConfig{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"email_domain": item.EmailDomain,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.School{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(id int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByName(name string) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *Repository) GetInfoByID(id int64) (*model.SchoolInfo, error) {
	result := &model.SchoolInfo{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetConfigByID(id int64) (*model.SchoolConfig, error) {
	result := &model.SchoolConfig{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.School, err error) {
	result = make([]model.School, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE CAST(schools.id AS TEXT) = '%s' OR schools.name ILIKE '%s' OR CAST(schools.type AS TEXT) ILIKE '%s' OR school_infos.full_name ILIKE '%s' OR school_infos.slogan ILIKE '%s' OR school_infos.founder ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	newFilter := filter
	newFilter.OrderBy = "schools." + newFilter.OrderBy
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT schools.created_at, schools.updated_at, schools.id, schools.name, schools.type, schools.created_at, schools.updated_at"+
				", schools.school_config_id, schools.school_info_id FROM schools "+
				"LEFT JOIN school_infos ON schools.school_info_id = school_infos.id",
			where,
			pagination,
			newFilter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
