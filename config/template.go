package config

import (
	"github.com/4kpros/go-api/common/helpers"
	"go.uber.org/zap"
)

type Template struct {
	Redoc     string
	Scalar    string
	Stoplight string
	Swagger   string
}

var AppTemplate = &Template{}

func LoadTemplates() (err error) {
	// Redoc
	AppTemplate.Redoc, err = helpers.ReadFileContentToString("templates/openapi/redoc.html")
	if err != nil {
		helpers.Logger.Warn(
			"Failed to load OpenAPI Redoc template",
			zap.String("Error", err.Error()),
		)
		return
	}

	// Scalar
	AppTemplate.Scalar, err = helpers.ReadFileContentToString("templates/openapi/scalar.html")
	if err != nil {
		helpers.Logger.Warn(
			"Failed to load OpenAPI Scalar template",
			zap.String("Error", err.Error()),
		)
		return
	}

	// Stoplight
	AppTemplate.Stoplight, err = helpers.ReadFileContentToString("templates/openapi/stoplight.html")
	if err != nil {
		helpers.Logger.Warn(
			"Failed to load OpenAPI Stoplight template",
			zap.String("Error", err.Error()),
		)
		return
	}

	// Swagger
	AppTemplate.Swagger, err = helpers.ReadFileContentToString("templates/openapi/swagger.html")
	if err != nil {
		helpers.Logger.Warn(
			"Failed to load OpenAPI Swagger template",
			zap.String("Error", err.Error()),
		)
		return
	}
	return
}
