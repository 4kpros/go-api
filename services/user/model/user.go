package model

import (
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user/data"
	"gorm.io/gorm"
)

type User struct {
	types.BaseGormModel
	Email       string `gorm:"default:null"`
	PhoneNumber uint64 `gorm:"default:null"`
	RoleId      int64  `gorm:"default:null"`
	Password    string `gorm:"default:null"`

	SignInMethod   string     `gorm:"default:null"`
	Provider       string     `gorm:"default:null"`
	ProviderUserId string     `gorm:"default:null"`
	IsActivated    bool       `gorm:"default:null"`
	ActivatedAt    *time.Time `gorm:"default:null"`

	UserInfo   UserInfo `gorm:"default:null;foreignKey:UserInfoId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	UserInfoId int64    `gorm:"default:null"`

	UserMfa   UserMfa `gorm:"default:null;foreignKey:UserMfaId;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
	UserMfaId int64   `gorm:"default:null"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Password, err = utils.EncodeArgon2id(user.Password)
	return
}

func (user *User) BeforeUpdate(db *gorm.DB) (err error) {
	user.Password, err = utils.EncodeArgon2id(user.Password)
	return
}
func (user *User) FromGoogleUser(googleUser *types.GoogleUserProfileResponse) {
	user.ProviderUserId = googleUser.ID
	user.Email = googleUser.Email
	user.UserInfo = UserInfo{
		UserName:  googleUser.FullName,
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
		Language:  googleUser.Language,
		Image:     googleUser.Picture,
	}
}
func (user *User) FromFacebookUser(facebookUser *types.FacebookUserProfileResponse) {
	user.ProviderUserId = facebookUser.ID
	user.Email = facebookUser.Email
	user.UserInfo = UserInfo{
		UserName:  facebookUser.FullName,
		FirstName: facebookUser.FirstName,
		LastName:  facebookUser.LastName,
		// Language:  facebookUser.Languages,
		Image: facebookUser.PictureSmall.Data.Url,
	}
}

func (user *User) ToResponse() *data.UserResponse {
	resp := &data.UserResponse{}
	resp.ID = user.ID
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
	resp.DeletedAt = user.DeletedAt

	resp.Email = user.Email
	resp.PhoneNumber = user.PhoneNumber
	resp.RoleId = user.RoleId

	resp.SignInMethod = user.SignInMethod
	resp.Provider = user.Provider
	resp.ProviderUserId = user.ProviderUserId
	resp.IsActivated = user.IsActivated
	resp.ActivatedAt = user.ActivatedAt

	resp.UserInfo = *user.UserInfo.ToResponse()
	resp.UserMfa = *user.UserMfa.ToResponse()
	return resp
}

func ToResponseList(userList []User) []data.UserResponse {
	resp := make([]data.UserResponse, len(userList))
	for index, user := range userList {
		resp[index] = *user.ToResponse()
	}
	return resp
}
