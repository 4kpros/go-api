package profile

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/constants"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/services/user"
	"github.com/4kpros/go-api/services/user/data"
	"github.com/4kpros/go-api/services/user/model"
)

type ProfileService struct {
	Repository *user.UserRepository
}

func NewProfileService(repository *user.UserRepository) *ProfileService {
	return &ProfileService{Repository: repository}
}

// Update profile
func (service *ProfileService) UpdateProfile(id int64, data *data.UpdateProfileRequest) (result *model.User, errCode int, err error) {
	result, err = service.Repository.UpdateProfile(id, &model.User{
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Password:    data.Password,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update user from database")
	}
	return
}

// Update profile info
func (service *ProfileService) UpdateProfileInfo(id int64, data *data.UpdateProfileInfoRequest) (result *model.UserInfo, errCode int, err error) {
	result, err = service.Repository.UpdateProfileInfo(id, &model.UserInfo{
		UserName:  data.UserName,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Address:   data.Address,
		Image:     data.Image,
		Language:  data.Language,
	})

	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update profile info from database")
	}
	return
}

// Update profile MFA
func (service *ProfileService) UpdateProfileMfa(id int64, data *data.UpdateProfileMfaRequest) (result *model.UserMfa, errCode int, err error) {
	if !utils.IsMfaMethodValid(data.Method) {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("Invalid MFA method! Please enter valid method.")
		return
	}
	result, err = service.Repository.UpdateProfileMfa(id, data.Method, data.Value)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("update profile MFA from database")
	}
	return
}

// Delete user account
func (service *ProfileService) Delete(id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
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

// Return profile info
func (service *ProfileService) GetById(id int64) (user *model.User, errCode int, err error) {
	user, err = service.Repository.GetById(id)
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
