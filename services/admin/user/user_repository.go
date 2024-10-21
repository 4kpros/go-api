package user

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/services/admin/user/model"
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

func (repository *Repository) UpdateUser(id int64, user *model.User) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"role_id":      user.RoleId,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.User{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByEmail(email string) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Where(
		"sign_in_method = ?", constants.AuthLoginMethodDefault,
	).Where(
		"email = ?", email,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetByPhoneNumber(phoneNumber uint64) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Where(
		"sign_in_method = ?", constants.AuthLoginMethodDefault,
	).Where(
		"phone_number = ?", phoneNumber,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetByProvider(provider string, providerUserId string) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Where(
		"sign_in_method = ?", constants.AuthLoginMethodProvider,
	).Where(
		"provider = ?", provider,
	).Where(
		"provider_user_id = ?", providerUserId,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	var result []model.User
	return result, repository.Db.Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
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
func (repository *Repository) UpdateUserPassword(id int64, password string) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Model(result).Where("id = ?", id).Update("password", password).Error
}

func (repository *Repository) UpdateUserActivation(id int64, user *model.User) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"user_info_id": user.UserInfoId,
			"user_mfa_id":  user.UserMfaId,
			"is_activated": user.IsActivated,
			"activated_at": user.ActivatedAt,
		},
	).Error
}

// ----------------- Profile service -----------------
func (repository *Repository) UpdateEmail(id int64, email string) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"email": email,
		},
	).Error
}
func (repository *Repository) UpdatePhoneNumber(id int64, phoneNumber uint64) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"phone_number": phoneNumber,
		},
	).Error
}
func (repository *Repository) UpdatePassword(id int64, password string) (*model.User, error) {
	var result *model.User
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"password": password,
		},
	).Error
}

func (repository *Repository) UpdateProfileInfo(id int64, userInfo *model.UserInfo) (*model.UserInfo, error) {
	var result *model.UserInfo
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"user_name":  userInfo.UserName,
			"first_name": userInfo.FirstName,
			"last_name":  userInfo.LastName,
			"address":    userInfo.Address,
			"image":      userInfo.Image,
			"language":   userInfo.Language,
		},
	).Error
}

func (repository *Repository) UpdateProfileMfa(id int64, column string, value bool) (*model.UserMfa, error) {
	var result *model.UserMfa
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"" + column: value,
		},
	).Error
}
