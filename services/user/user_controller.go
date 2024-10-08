package user

type UserController struct {
	Service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{Service: service}
}

// func (controller *UserController) CreateWithEmail(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var user = &model.User{}
// 	user.Email = input.Email
// 	user.Role = input.Role

// 	// Execute the service
// 	password, _, err := controller.Service.CreateWithEmail(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &response.CreateWithEmailResponse{
// 		Email:    user.Email,
// 		Role:     user.Role,
// 		Password: password,
// 	}, nil
// }

// func (controller *UserController) CreateWithPhoneNumber(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var reqData = &request.CreateWithPhoneNumberRequest{}
// 	ctx.Bind(reqData)
// 	var user = &model.User{}
// 	user.PhoneNumber = reqData.PhoneNumber
// 	user.Role = reqData.Role

// 	// Execute the service
// 	password, errCode, err := controller.Service.CreateWithPhoneNumber(user)
// 	if err != nil {
// 		ctx.AbortWithError(errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, response.CreateWithPhoneNumberResponse{
// 		PhoneNumber: reqData.PhoneNumber,
// 		Role:        reqData.Role,
// 		Password:    password,
// 	})
// }

// func (controller *UserController) UpdateUser(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var user = &model.User{}
// 	ctx.Bind(user)

// 	// Execute the service
// 	errCode, err := controller.Service.UpdateUser(user)
// 	if err != nil {
// 		ctx.AbortWithError(errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, user)
// }

// func (controller *UserController) UpdateUserInfo(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	var userInfo = &model.UserInfo{}
// 	ctx.Bind(userInfo)

// 	// Execute the service
// 	errCode, err := controller.Service.UpdateUserInfo(userInfo)
// 	if err != nil {
// 		ctx.AbortWithError(errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, userInfo)
// }

// func (controller *UserController) Delete(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req header
// 	id := ctx.Param("id")

// 	// Execute the service
// 	user, errCode, err := controller.Service.Delete(id)
// 	if err != nil {
// 		ctx.AbortWithError(errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, user)
// }

// func (controller *UserController) FindById(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req header
// 	id := ctx.Param("id")

// 	// Execute the service
// 	user, errCode, err := controller.Service.FindById(id)
// 	if err != nil {
// 		ctx.AbortWithError(errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, user)
// }

// func (controller *UserController) FindUserInfoById(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req header
// 	id := ctx.Param("id")

// 	// Execute the service
// 	user, errCode, err := controller.Service.FindUserInfoById(id)
// 	if err != nil {
// 		ctx.AbortWithError(errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, user)
// }

// func (controller *UserController) FindAll(ctx huma.Context, next func(huma.Context)) {
// 	// Get data of req body
// 	pagination, filter := utils.GetPaginationFiltersFromQuery(c)

// 	// Execute the service
// 	users, errCode, err := controller.Service.FindAll(filter, pagination)
// 	if err != nil {
// 		ctx.AbortWithError(errCode, err)
// 		return
// 	}

// 	// Return the response
// 	ctx.JSON(http.StatusOK, types.SuccessPaginatedResponse{
// 		Data:       users,
// 		Filter:     filter,
// 		Pagination: pagination,
// 	})
// }
