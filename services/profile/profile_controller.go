package profile

import (
	"github.com/4kpros/go-api/services/user/data"
	"github.com/4kpros/go-api/services/user/model"
)

type ProfileController struct {
	Service *ProfileService
}

func NewProfileController(service *ProfileService) *ProfileController {
	return &ProfileController{Service: service}
}

func (controller *ProfileController) UpdateProfile(id int64, data *data.UpdateProfileRequest) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfile(id, data)
	return
}

func (controller *ProfileController) UpdateProfileInfo(id int64, data *data.UpdateProfileInfoRequest) (result *model.UserInfo, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileInfo(id, data)
	return
}

func (controller *ProfileController) UpdateProfileMfa(id int64, data *data.UpdateProfileMfaRequest) (result *model.UserMfa, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileMfa(id, data)
	return
}

func (controller *ProfileController) Delete(id int64) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(id)
	return
}

func (controller *ProfileController) GetById(id int64) (result *model.User, errCode int, err error) {
	user, errCode, err := controller.Service.GetById(id)
	if err != nil {
		return
	}
	result = user
	return
}
