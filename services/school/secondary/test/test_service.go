package test

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/secondary/test/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new test
func (service *Service) Create(inputJwtToken *types.JwtToken, test *model.Test) (result *model.Test, errCode int, err error) {
	// Check if test already exists
	foundTest, err := service.Repository.GetByObject(&model.Test{
		SchoolID:  test.SchoolID,
		SubjectID: test.SubjectID,
		Type:      test.Type,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get test by name from database")
		return
	}
	if foundTest != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("test")
		return
	}

	// Insert test
	result, err = service.Repository.Create(test)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create test from database")
		return
	}
	return
}

// Update test
func (service *Service) Update(inputJwtToken *types.JwtToken, testID int64, test *model.Test) (result *model.Test, errCode int, err error) {
	// Check if test already exists
	foundTestByID, err := service.Repository.GetById(testID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get test by name from database")
		return
	}
	if foundTestByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Test")
		return
	}
	foundTest, err := service.Repository.GetByObject(&model.Test{
		SchoolID:  test.SchoolID,
		SubjectID: test.SubjectID,
		Type:      test.Type,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get test by name from database")
		return
	}
	if foundTest != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("test")
		return
	}

	// Update test
	result, err = service.Repository.Update(testID, inputJwtToken.UserID, test)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update test from database")
		return
	}
	return
}

// Delete test with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete test from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Test")
		return
	}
	return
}

// Get Returns test with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, testID int64) (test *model.Test, errCode int, err error) {
	test, err = service.Repository.GetById(testID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get test by id from database")
		return
	}
	if test == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Test")
		return
	}
	return
}

// GetAll Returns all tests with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (testList []model.Test, errCode int, err error) {
	testList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get tests from database")
	}
	return
}
