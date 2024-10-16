package profile

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/user"
	"api/services/user/model"
)

type ProfileService struct {
	Repository *user.UserRepository
}

func NewProfileService(repository *user.UserRepository) *ProfileService {
	return &ProfileService{Repository: repository}
}

// Update profile
func (service *ProfileService) UpdateProfile(jwtToken *types.JwtToken, user *model.User) (result *model.User, errCode int, err error) {
	result, err = service.Repository.UpdateProfile(jwtToken.UserId, user)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update user from database")
	}
	return
}

// Update profile info
func (service *ProfileService) UpdateProfileInfo(jwtToken *types.JwtToken, userInfo *model.UserInfo) (result *model.UserInfo, errCode int, err error) {
	result, err = service.Repository.UpdateProfileInfo(jwtToken.UserId, userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update profile info from database")
	}
	return
}

// Update profile MFA
func (service *ProfileService) UpdateProfileMfa(jwtToken *types.JwtToken, method string, value bool) (result *model.UserMfa, errCode int, err error) {
	if !utils.IsMfaMethodValid(method) {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid MFA method! Please enter valid method.")
		return
	}
	result, err = service.Repository.UpdateProfileMfa(jwtToken.UserId, method, value)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update profile MFA from database")
	}
	return
}

// Delete user account
func (service *ProfileService) DeleteProfile(jwtToken *types.JwtToken) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(jwtToken.UserId)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("delete account from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("User")
		return
	}
	return
}

// Return profile information
func (service *ProfileService) GetProfile(jwtToken *types.JwtToken) (user *model.User, errCode int, err error) {
	user, err = service.Repository.GetById(jwtToken.UserId)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("get profile from database")
		return
	}
	if user == nil {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("User")
		return
	}
	return
}
