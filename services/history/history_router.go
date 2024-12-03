package history

import (
	"api/common/constants"
	"api/common/types"
	"api/services/history/data"

	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/history",
		Tag:   []string{"History"},
	}
	const tableName = "history"

	// Get all history
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-history-list",
			Summary:     "Get history list" + " (" + constants.FeatureAdminLabel + ")",
			Description: "Get history list with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
						tableName,                // Table name
						constants.PermissionRead, // Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden},
		},
		func(
			ctx context.Context,
			input *struct {
				types.Filter
				types.PaginationRequest
			},
		) (*struct {
			Body data.HistoryList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.HistoryList
			}{Body: *result}, nil
		},
	)
}
