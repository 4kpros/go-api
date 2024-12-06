package student

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/student/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new student
func (service *Service) Create(inputJwtToken *types.JwtToken, student *model.Student) (result *model.Student, errCode int, err error) {
	// Check if student already exists
	foundStudent, err := service.Repository.GetByObject(&model.Student{
		SchoolID: student.SchoolID,
		UserID:   student.UserID,
		LevelID:  student.LevelID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by name from database")
		return
	}
	if foundStudent != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("student")
		return
	}

	// Insert student
	result, err = service.Repository.Create(student)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create student from database")
		return
	}
	return
}

// Update student
func (service *Service) Update(inputJwtToken *types.JwtToken, studentID int64, student *model.Student) (result *model.Student, errCode int, err error) {
	// Check if student already exists
	foundStudentByID, err := service.Repository.GetById(studentID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by name from database")
		return
	}
	if foundStudentByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student")
		return
	}
	foundStudent, err := service.Repository.GetByObject(&model.Student{
		SchoolID: foundStudentByID.SchoolID,
		UserID:   foundStudentByID.UserID,
		LevelID:  student.LevelID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by name from database")
		return
	}
	if foundStudent != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("student")
		return
	}

	// Update student
	result, err = service.Repository.Update(studentID, inputJwtToken.UserID, student)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update student from database")
		return
	}
	return
}

// Delete student with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, studentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(studentID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete student from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student")
		return
	}
	return
}

// Get Returns student with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, studentID int64) (student *model.Student, errCode int, err error) {
	student, err = service.Repository.GetById(studentID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student by id from database")
		return
	}
	if student == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Student")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (studentList []model.Student, errCode int, err error) {
	studentList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
