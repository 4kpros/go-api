package model

import (
	"api/services/admin/role/model"
	"time"

	"gorm.io/gorm"

	"api/common/types"
	"api/common/utils/security"
	"api/services/admin/user/data"
)

type User struct {
	types.BaseGormModel
	Email       string `gorm:"default:null"`
	PhoneNumber uint64 `gorm:"default:null"`
	Password    string `gorm:"default:null"`

	LoginMethod    string     `gorm:"default:null"`
	Provider       string     `gorm:"default:null"`
	ProviderUserId string     `gorm:"default:null"`
	IsActivated    bool       `gorm:"default:null"`
	ActivatedAt    *time.Time `gorm:"default:null"`

	Role   *model.Role `gorm:"default:null;foreignKey:RoleId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	RoleId int64       `gorm:"default:null"`

	UserInfo   *UserInfo `gorm:"default:null;foreignKey:UserInfoId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	UserInfoId int64     `gorm:"default:null"`

	UserMfa   *UserMfa `gorm:"default:null;foreignKey:UserMfaId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	UserMfaId int64    `gorm:"default:null"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Password, err = security.EncodeArgon2id(user.Password)
	return
}

func (user *User) BeforeUpdate(db *gorm.DB) (err error) {
	user.Password, err = security.EncodeArgon2id(user.Password)
	return
}

func (user *User) ToResponse() *data.UserResponse {
	resp := &data.UserResponse{
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,

		LoginMethod:    user.LoginMethod,
		Provider:       user.Provider,
		ProviderUserId: user.ProviderUserId,
		IsActivated:    user.IsActivated,
		ActivatedAt:    user.ActivatedAt,

		Role: &data.UserRoleResponse{
			ID:          user.Role.ID,
			Name:        user.Role.Name,
			Description: user.Role.Description,
		},
		UserInfo: user.UserInfo.ToResponse(),
		UserMfa:  user.UserMfa.ToResponse(),
	}
	resp.ID = user.ID
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
	resp.DeletedAt = user.DeletedAt
	return resp
}

func (user *User) FromGoogleUser(googleUser *types.GoogleUserProfileResponse) {
	user.ProviderUserId = googleUser.ID
	user.Email = googleUser.Email
	user.UserInfo = &UserInfo{
		UserName:  googleUser.FullName,
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
		// Language:  googleUser.Language,
		Image: googleUser.Picture,
	}
}
func (user *User) FromFacebookUser(facebookUser *types.FacebookUserProfileResponse) {
	user.ProviderUserId = facebookUser.ID
	user.Email = facebookUser.Email
	user.UserInfo = &UserInfo{
		UserName:  facebookUser.FullName,
		FirstName: facebookUser.FirstName,
		LastName:  facebookUser.LastName,
		// Language:  facebookUser.Languages,
		Image: facebookUser.PictureSmall.Data.Url,
	}
}

func ToResponseList(userList []User) []data.UserResponse {
	resp := make([]data.UserResponse, len(userList))
	for index, user := range userList {
		resp[index] = *user.ToResponse()
	}
	return resp
}
