package admin

import (
	"context"

	"api/services/admin/data"
	"api/services/user/model"
)

type AdminController struct {
	Service *AdminService
}

func NewAdminController(service *AdminService) *AdminController {
	return &AdminController{Service: service}
}

func (controller *AdminController) Create(
	ctx *context.Context,
	input *struct {
		Body data.CreateAdminRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		input.Body.Token,
		&model.User{
			Email:    input.Body.Email,
			Password: input.Body.Password,
		},
	)
	return
}
