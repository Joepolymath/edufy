package school

import "time"

type SetUpSchoolDto struct {
	SchoolName  string  `form:"schoolName" validate:"required,min=2,max=50"`
	Email       string  `form:"email" validate:"required,min=2,max=50"`
	SchoolType  string  `form:"schoolType" validate:"required,min=2,max=50,validSchoolType"`
	Logo        *string `form:"logo"`
	Document    *string `form:"document"`
	Address     string  `form:"address" validate:"required,min=5,max=100"`
	PhoneNumber string  `form:"phoneNumber" validate:"required,min=10,max=15,numeric"`
	BrandColor  *string `form:"brandColor" validate:"required,max=250"`
}

type AddSchoolBranchDto struct {
	BranchName  string `json:"branchName" validate:"required,min=2,max=50"`
	Email       string `json:"email" validate:"required,min=2,max=50"`
	SchoolType  string `json:"schoolType" validate:"required,min=2,max=50,validSchoolType"`
	Address     string `json:"address" validate:"required,min=5,max=100"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=15,numeric"`
}

type CreateSessionDto struct {
	SessionName  string    `json:"sessionName"`
	SessionType  string    `json:"sessionType"`
	SessionStart time.Time `json:"sessionStart"`
	SessionEnd   time.Time `json:"sessionEnd"`
}
