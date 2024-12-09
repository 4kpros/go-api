package user

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/services/user/user/model"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(user *model.User) (*model.User, error) {
	result := *user
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) AssignRole(userID int64, roleID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", userID).Updates(
		map[string]interface{}{
			"role_id": roleID,
		},
	).Error
}

func (repository *Repository) Update(userID int64, user *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", userID).Updates(
		map[string]interface{}{
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
		},
	).Error
}

func (repository *Repository) Delete(userID int64) (int64, error) {
	result := repository.Db.Where("id = ?", userID).Delete(&model.User{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteRole(userID int64, roleID int64) (int64, error) {
	result := repository.Db.Model(&model.User{}).Where("id = ?", userID).Where("role_id = ?", roleID).Updates(
		map[string]interface{}{
			"role_id": nil,
		},
	)
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(userID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where("id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByEmail(email string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where(
		"login_method = ?", constants.AuthLoginMethodDefault,
	).Where(
		"email = ?", email,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetByPhoneNumber(phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where(
		"login_method = ?", constants.AuthLoginMethodDefault,
	).Where(
		"phone_number = ?", phoneNumber,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetByProvider(provider string, providerUserID string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Where(
		"login_method = ?", constants.AuthLoginMethodProvider,
	).Where(
		"provider = ?", provider,
	).Where(
		"provider_user_id = ?", providerUserID,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	var result []model.User
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE name ILIKE '%s' OR feature ILIKE '%s' OR description ILIKE '%s'",
			filter.Search,
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM users",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error
}

// ----------------- Authentication service -----------------

func (repository *Repository) CreateUserInfo(userInfo *model.UserInfo) (*model.UserInfo, error) {
	result := *userInfo
	return &result, repository.Db.Create(&result).Error
}
func (repository *Repository) CreateUserMfa(userMfa *model.UserMfa) (*model.UserMfa, error) {
	result := *userMfa
	return &result, repository.Db.Create(&result).Error
}
func (repository *Repository) UpdateUserPassword(userID int64, password string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", userID).Update("password", password).Error
}

func (repository *Repository) UpdateUserActivation(userID int64, user *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", userID).Updates(
		map[string]interface{}{
			"is_activated": user.IsActivated,
			"activated_at": user.ActivatedAt,
		},
	).Error
}

// ----------------- Profile service -----------------

func (repository *Repository) UpdateEmail(userID int64, email string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", userID).Updates(
		map[string]interface{}{
			"email": email,
		},
	).Error
}
func (repository *Repository) UpdatePhoneNumber(userID int64, phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", userID).Updates(
		map[string]interface{}{
			"phone_number": phoneNumber,
		},
	).Error
}
func (repository *Repository) UpdatePassword(userID int64, password string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Model(result).Where("id = ?", userID).Updates(
		map[string]interface{}{
			"password": password,
		},
	).Error
}

func (repository *Repository) UpdateProfileInfo(userInfoID int64, userInfo *model.UserInfo) (*model.UserInfo, error) {
	result := &model.UserInfo{}
	return result, repository.Db.Model(result).Where("id = ?", userInfoID).Updates(
		map[string]interface{}{
			"username":   userInfo.Username,
			"first_name": userInfo.FirstName,
			"last_name":  userInfo.LastName,
			"address":    userInfo.Address,
			"image":      userInfo.Image,
			"language":   userInfo.Language,
		},
	).Error
}

func (repository *Repository) UpdateProfileMfa(userMfaID int64, column string, value bool) (*model.UserMfa, error) {
	result := &model.UserMfa{}
	return result, repository.Db.Model(result).Where("id = ?", userMfaID).Updates(
		map[string]interface{}{
			"" + column: value,
		},
	).Error
}
