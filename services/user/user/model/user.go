package model

import (
	"time"

	"gorm.io/gorm"

	"api/common/types"
	"api/common/utils/security"
	"api/services/user/role/model"
	"api/services/user/user/data"
)

type User struct {
	types.BaseGormModel
	Email       string `gorm:"default:null"`
	PhoneNumber uint64 `gorm:"default:null"`
	Password    string `gorm:"default:null"`

	LoginMethod    string     `gorm:"default:null"`
	Provider       string     `gorm:"default:null"`
	ProviderUserID string     `gorm:"default:null"`
	IsActivated    bool       `gorm:"default:null"`
	ActivatedAt    *time.Time `gorm:"default:null"`

	RoleID int64       `gorm:"default:null"`
	Role   *model.Role `gorm:"default:null;foreignKey:RoleID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserInfoID int64     `gorm:"default:null"`
	UserInfo   *UserInfo `gorm:"default:null;foreignKey:UserInfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserMfaID int64    `gorm:"default:null"`
	UserMfa   *UserMfa `gorm:"default:null;foreignKey:UserMfaID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *User) BeforeCreate(db *gorm.DB) (err error) {
	item.Password, err = security.EncodeArgon2id(item.Password)
	return
}

func (item *User) BeforeUpdate(db *gorm.DB) (err error) {
	item.Password, err = security.EncodeArgon2id(item.Password)
	return
}

func (item *User) ToResponse() *data.UserResponse {
	resp := &data.UserResponse{}
	if item == nil {
		return resp
	}
	resp.Email = item.Email
	resp.PhoneNumber = item.PhoneNumber
	resp.LoginMethod = item.LoginMethod
	resp.Provider = item.Provider
	resp.ProviderUserID = item.ProviderUserID
	resp.IsActivated = item.IsActivated
	resp.ActivatedAt = item.ActivatedAt

	resp.Role = item.Role.ToResponse()
	resp.UserInfo = item.UserInfo.ToResponse()
	resp.UserMfa = item.UserMfa.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *User) FromGoogleUser(googleUser *types.GoogleUserProfileResponse) {
	item.ProviderUserID = googleUser.ID
	item.Email = googleUser.Email
	item.UserInfo = &UserInfo{
		Username:  googleUser.FullName,
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
		// Language:  googleUser.Language,
		Image: googleUser.Picture,
	}
}
func (item *User) FromFacebookUser(facebookUser *types.FacebookUserProfileResponse) {
	item.ProviderUserID = facebookUser.ID
	item.Email = facebookUser.Email
	item.UserInfo = &UserInfo{
		Username:  facebookUser.FullName,
		FirstName: facebookUser.FirstName,
		LastName:  facebookUser.LastName,
		// Language:  facebookUser.Languages,
		Image: facebookUser.PictureSmall.Data.Url,
	}
}

func ToResponseList(itemList []User) []data.UserResponse {
	resp := make([]data.UserResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
