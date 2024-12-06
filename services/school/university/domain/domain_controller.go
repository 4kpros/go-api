package domain

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/domain/data"
	"api/services/school/university/domain/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Create(
	ctx *context.Context,
	input *struct {
		Body data.CreateDomainRequest
	},
) (result *model.Domain, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Domain{
			SchoolID:     input.Body.SchoolID,
			DepartmentID: input.Body.DepartmentID,
			Name:         input.Body.Name,
			Description:  input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.DomainID
		Body data.UpdateDomainRequest
	},
) (result *model.Domain, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Domain{
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.DomainID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.DomainID
	},
) (result *model.Domain, errCode int, err error) {
	domain, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = domain
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.DomainResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	domainList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination)
	if err != nil {
		return
	}
	result = &data.DomainResponseList{
		Data: model.ToResponseList(domainList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
