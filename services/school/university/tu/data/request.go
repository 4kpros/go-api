package data

type TeachingUnitID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teaching unit id" example:"1"`
}

type CreateTeachingUnitRequest struct {
	SchoolID     int64  `json:"schoolID" required:"false" doc:"School id" example:"1"`
	DomainID     int64  `json:"domainID" required:"false" doc:"Domain id" example:"1"`
	LevelID      int64  `json:"levelID" required:"false" doc:"Level id" example:"1"`
	Name         string `json:"name" required:"false" doc:"Name" example:"MATH110"`
	Description  string `json:"description" required:"false" doc:"Description" example:""`
	Credit       int    `json:"credit" required:"false" doc:"Credit" example:"1"`
	Semester     int    `json:"semester" required:"true" doc:"Semester" example:"1"`
	Program      string `json:"Program" required:"false" doc:"Program" example:""`
	Requirements string `json:"Requirements" required:"false" doc:"Requirements" example:""`
}

type TeachingUnitProfessorRequest struct {
	UserID int64 `json:"userID" required:"true" doc:"User id" example:"1"`
}

type UpdateTeachingUnitRequest struct {
	Name         string `json:"name" required:"false" doc:"Name" example:"MATH110"`
	Description  string `json:"description" required:"false" doc:"Description" example:""`
	Credit       int    `json:"credit" required:"false" doc:"Credit" example:"1"`
	Semester     int    `json:"semester" required:"true" doc:"Semester" example:"1"`
	Program      string `json:"Program" required:"false" doc:"Program" example:""`
	Requirements string `json:"Requirements" required:"false" doc:"Requirements" example:""`
}
