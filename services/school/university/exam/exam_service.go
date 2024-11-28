package exam

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/exam/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new exam
func (service *Service) Create(inputJwtToken *types.JwtToken, exam *model.Exam) (result *model.Exam, errCode int, err error) {
	// Check if exam already exists
	foundExam, err := service.Repository.GetByObject(&model.Exam{
		SchoolID:       exam.SchoolID,
		TeachingUnitID: exam.TeachingUnitID,
		Type:           exam.Type,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundExam != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("exam")
		return
	}

	// Insert exam
	result, err = service.Repository.Create(exam)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create exam from database")
		return
	}
	return
}

// Update exam
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, exam *model.Exam) (result *model.Exam, errCode int, err error) {
	// Check if exam already exists
	foundExamByID, err := service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundExamByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	foundExam, err := service.Repository.GetByObject(&model.Exam{
		SchoolID:       exam.SchoolID,
		TeachingUnitID: exam.TeachingUnitID,
		Type:           exam.Type,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by name from database")
		return
	}
	if foundExam != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("exam")
		return
	}

	// Update exam
	result, err = service.Repository.Update(id, inputJwtToken.UserID, exam)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update exam from database")
		return
	}
	return
}

// Delete exam with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete exam from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	return
}

// Get Returns exam with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (exam *model.Exam, errCode int, err error) {
	exam, err = service.Repository.GetById(id, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exam by id from database")
		return
	}
	if exam == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Exam")
		return
	}
	return
}

// GetAll Returns all exams with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (examList []model.Exam, errCode int, err error) {
	examList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get exams from database")
	}
	return
}
