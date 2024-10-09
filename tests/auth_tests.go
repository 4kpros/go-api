package tests

import (
	"testing"

	"github.com/4kpros/go-api/services/auth"
)

func AuthTests(t *testing.T) {
	// TODO generate database mock
	repo := auth.NewAuthRepositoryImpl(nil)
	service := auth.NewAuthServiceImpl(repo)
	controller := auth.NewAuthController(service)

	if repo == nil || service == nil || controller == nil {
		return
	}

	// TODO add tests
}
