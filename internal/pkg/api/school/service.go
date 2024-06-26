package school

import (
	"Learnium/internal/pkg/models"
	"context"
	"time"
	// "github.com/patrickmn/go-cache"
)

// describe service to be injected by school controller to accomplish school related tasks
type ISchoolService interface {
	AddNewSchool(ctx context.Context, ownerId, name, email, schoolType, mobile,
		address, brandColor, logo, document string) (models.School, error)
	CreateSession(ctx context.Context, schoolId, name, sessionType string,
		sessionStart, sessionEnd time.Time) (models.Session, error)
	AddSchoolBranch(ctx context.Context, parentSchool models.School, name, email, schoolType,
		address, mobile string) (models.School, error)
	// GetSchoolByOwnerID(ctx context.Context, id string) (models.School, error)
	GetParentSchoolByOwnerID(ctx context.Context, id string) (models.School, error)
	GetSchoolsByOwnerID(ctx context.Context, id string) (*SchoolsData, error)
	CreateDefaultSchoolRoles(ctx context.Context, schoolId string) error
	GetSchoolRolesByType(ctx context.Context, schoolId string, roleType models.RoleType) ([]models.SchoolRole, error)
}

type SchoolsData struct {
	ParentSchool  *models.School
	BranchSchools []models.School
}

type SchoolService struct {
	repository ISchoolRepository
}

// create new school service instance
func NewSchoolService(repository ISchoolRepository) ISchoolService {
	Srv := &SchoolService{
		repository,
	}
	service := ISchoolService(Srv)
	return service
}

func (srv *SchoolService) AddNewSchool(ctx context.Context, ownerId, name, email, schoolType, mobile,
	address, brandColor, logo, document string) (models.School, error) {
	sch := &models.School{
		Mail: email, Name: name, Type: models.SchoolType(schoolType),
		Phone: mobile, BrandColor: &brandColor, Logo: &logo,
		RegistrationDocument: &document, Address: &address,
		OwnerID: ownerId, IsParentSchool: true, IsBranchSchool: false,
	}
	return srv.repository.CreateSchool(ctx, *sch)

}

func (srv *SchoolService) CreateSession(ctx context.Context, schoolId, name, sessionType string,
	sessionStart, sessionEnd time.Time) (models.Session, error) {
	sess := &models.Session{
		SchoolID: &schoolId, Name: name, Type: &sessionType,
		EndDate: &sessionEnd, StartDate: &sessionStart,
	}
	return srv.repository.CreateSession(ctx, *sess)

}

func (srv *SchoolService) CreateDefaultSchoolRoles(ctx context.Context, schoolId string) error {
	defaultSchoolRoles := []models.SchoolRole{
		{Name: "School Administrator", RoleType: models.ADMIN,
			Description: &models.SchoolAdminDesc, SchoolID: &schoolId},

		{Name: "Head Of Department", RoleType: models.TEACHING_STAFF,
			SchoolID: &schoolId, Description: &models.HODDesc},
		{Name: "Class Teacher", RoleType: models.TEACHING_STAFF,
			Description: &models.CTDesc, SchoolID: &schoolId},
		{Name: "Subject Teacher", RoleType: models.TEACHING_STAFF,
			Description: &models.STDesc, SchoolID: &schoolId},
		{Name: "Curriculum Developer", RoleType: models.TEACHING_STAFF,
			Description: &models.CDevDesc, SchoolID: &schoolId},
		{Name: "Examinations Coordinator", RoleType: models.TEACHING_STAFF,
			Description: &models.ECDesc, SchoolID: &schoolId},

		{Name: "Accountant", RoleType: models.NON_TEACHING_STAFF,
			Description: &models.ACCDesc, SchoolID: &schoolId},
		{Name: "Counselor", RoleType: models.NON_TEACHING_STAFF,
			Description: &models.CounselDesc, SchoolID: &schoolId},
		{Name: "Admissions Officer", RoleType: models.NON_TEACHING_STAFF,
			Description: &models.AODesc, SchoolID: &schoolId},
		{Name: "Human Resources", RoleType: models.NON_TEACHING_STAFF,
			Description: &models.HRDesc, SchoolID: &schoolId},
		{Name: "Health and Wellnesss Coordinator", RoleType: models.NON_TEACHING_STAFF,
			Description: &models.HWCDesc, SchoolID: &schoolId},
	}

	return srv.repository.CreateSchoolRoles(ctx, defaultSchoolRoles)
}

func (srv *SchoolService) AddSchoolBranch(ctx context.Context, parentSchool models.School,
	name, email, schoolType, address, mobile string) (models.School, error) {

	branch := &models.School{
		Mail: email, Name: name, Type: models.SchoolType(schoolType),
		Phone: mobile, Address: &address, ParentSchoolID: &parentSchool.ID,
		Logo: parentSchool.Logo, RegistrationDocument: parentSchool.RegistrationDocument,
		BrandColor: parentSchool.BrandColor, IsParentSchool: false, IsBranchSchool: true,
		OwnerID: parentSchool.OwnerID,
	}
	return srv.repository.CreateSchool(ctx, *branch)

}

func (srv *SchoolService) GetParentSchoolByOwnerID(ctx context.Context, id string) (models.School, error) {
	return srv.repository.GetParentSchoolByOwnerID(ctx, id)

}

func (srv *SchoolService) GetSchoolsByOwnerID(ctx context.Context, id string) (*SchoolsData, error) {
	parentSchool, err := srv.repository.GetParentSchoolByOwnerID(ctx, id)
	if err != nil {
		return nil, err
	}
	branches, err := srv.repository.GetSchoolBranches(ctx, parentSchool.ID)
	if err != nil {
		return &SchoolsData{
			ParentSchool: &parentSchool,
		}, nil
	}
	return &SchoolsData{
		ParentSchool: &parentSchool, BranchSchools: branches,
	}, nil
}

func (srv *SchoolService) GetSchoolRolesByType(ctx context.Context, schoolId string, roleType models.RoleType) ([]models.SchoolRole, error) {
	conditions := map[string]interface{}{
		"role_type": roleType,
		"school_id": schoolId,
	}
	schoolRoles, err := srv.repository.FilterSchoolRole(ctx, conditions)
	if err != nil {
		return nil, err
	}
	return schoolRoles, nil
}
