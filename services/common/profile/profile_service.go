package profile

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/admin/user"
	"api/services/admin/user/model"
)

type Service struct {
	Repository *user.Repository
}

func NewProfileService(repository *user.Repository) *Service {
	return &Service{Repository: repository}
}

// UpdateProfileEmail Updates email
func (service *Service) UpdateProfileEmail(jwtToken *types.JwtToken, email string) (result *model.User, errCode int, err error) {
	result, err = service.Repository.UpdateEmail(jwtToken.UserId, email)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update user from database")
	}
	return
}

// UpdateProfilePhoneNumber Updates phone number
func (service *Service) UpdateProfilePhoneNumber(jwtToken *types.JwtToken, phoneNumber int64) (result *model.User, errCode int, err error) {
	result, err = service.Repository.UpdatePhoneNumber(jwtToken.UserId, phoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update user from database")
	}
	return
}

// UpdateProfilePassword Updates password
func (service *Service) UpdateProfilePassword(jwtToken *types.JwtToken, password string) (result *model.User, errCode int, err error) {
	result, err = service.Repository.UpdatePassword(jwtToken.UserId, password)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update user from database")
	}
	return
}

// UpdateProfileInfo Update profile info
func (service *Service) UpdateProfileInfo(jwtToken *types.JwtToken, userInfo *model.UserInfo) (result *model.UserInfo, errCode int, err error) {
	result, err = service.Repository.UpdateProfileInfo(jwtToken.UserId, userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update profile info from database")
	}
	return
}

// UpdateProfileMfa Update profile MFA
func (service *Service) UpdateProfileMfa(jwtToken *types.JwtToken, method string, value bool) (result *model.UserMfa, errCode int, err error) {
	if !utils.IsMfaMethodValid(method) {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid MFA method! Please enter valid method.")
		return
	}
	result, err = service.Repository.UpdateProfileMfa(jwtToken.UserId, method, value)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update profile MFA from database")
	}
	return
}

// DeleteProfile Delete user account
func (service *Service) DeleteProfile(jwtToken *types.JwtToken) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(jwtToken.UserId)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete account from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	return
}

// GetProfile Return profile information
func (service *Service) GetProfile(jwtToken *types.JwtToken) (user *model.User, errCode int, err error) {
	user, err = service.Repository.GetById(jwtToken.UserId)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get profile from database")
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	return
}
