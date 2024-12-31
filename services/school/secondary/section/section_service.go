package section

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/secondary/section/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

// Create new section
func (service *Service) Create(inputJwtToken *types.JwtToken, section *model.Section) (result *model.Section, errCode int, err error) {
	// Check if section already exists
	foundSection, err := service.Repository.GetByObject(&model.Section{
		SchoolID: section.SchoolID,
		Name:     section.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by name from database")
		return
	}
	if foundSection != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("section")
		return
	}

	// Insert section
	result, err = service.Repository.Create(section)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("create section from database")
		return
	}
	return
}

// Update section
func (service *Service) Update(inputJwtToken *types.JwtToken, sectionID int64, section *model.Section) (result *model.Section, errCode int, err error) {
	// Check if section already exists
	foundSectionByID, err := service.Repository.GetById(sectionID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by name from database")
		return
	}
	if foundSectionByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Section")
		return
	}
	foundSection, err := service.Repository.GetByObject(&model.Section{
		SchoolID: foundSectionByID.SchoolID,
		Name:     section.Name,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by name from database")
		return
	}
	if foundSection != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("section")
		return
	}

	// Update section
	result, err = service.Repository.Update(sectionID, inputJwtToken.UserID, section)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update section from database")
		return
	}
	return
}

// Delete section with matching id and return affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, sectionID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(sectionID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete section from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Section")
		return
	}
	return
}

// Get Returns section with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, sectionID int64) (section *model.Section, errCode int, err error) {
	section, err = service.Repository.GetById(sectionID, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get section by id from database")
		return
	}
	if section == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("Section")
		return
	}
	return
}

// GetAll Returns all faculties with support for search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (sectionList []model.Section, errCode int, err error) {
	sectionList, err = service.Repository.GetAll(filter, pagination, inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get faculties from database")
	}
	return
}
