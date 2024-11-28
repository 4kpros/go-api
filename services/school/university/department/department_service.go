package department

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/department/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new department
func (service *Service) Create(inputJwtToken *types.JwtToken, department *model.Department) (result *model.Department, errCode int, err error) {
	// Check if department already exists
	foundDepartment, err := service.Repository.GetByObject(&model.Department{
		FacultyID: department.FacultyID,
		Name:      department.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by name from database")
		return
	}
	if foundDepartment != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("department")
		return
	}

	// Insert department
	result, err = service.Repository.Create(department)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create department from database")
		return
	}
	return
}

// Update department
func (service *Service) Update(inputJwtToken *types.JwtToken, departmentID int64, department *model.Department) (result *model.Department, errCode int, err error) {
	// Check if department already exists
	foundDepartmentByID, err := service.Repository.GetById(departmentID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by name from database")
		return
	}
	if foundDepartmentByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	foundDepartment, err := service.Repository.GetByObject(&model.Department{
		FacultyID: foundDepartmentByID.FacultyID,
		Name:      department.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by name from database")
		return
	}
	if foundDepartment != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("department")
		return
	}

	// Update department
	result, err = service.Repository.Update(departmentID, inputJwtToken.UserID, department)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update department from database")
		return
	}
	return
}

// Delete department with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, departmentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(departmentID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete department from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	return
}

// Get Returns department with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, departmentID int64) (department *model.Department, errCode int, err error) {
	department, err = service.Repository.GetById(departmentID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get department by id from database")
		return
	}
	if department == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Department")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (departmentList []model.Department, errCode int, err error) {
	departmentList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
