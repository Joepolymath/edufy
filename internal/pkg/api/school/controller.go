package school

import (
	"Learnium/internal/pkg/adapters"
	"Learnium/internal/pkg/common"
	"Learnium/internal/pkg/models"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SchoolController struct {
	service         ISchoolService
	mailProvider    adapters.IEmailAdapter
	storageProvider adapters.IFileUploader
	serializer      common.IValidator
	logger          common.ILogger
}

// create new instance of School controller
func NewSchoolController(srv ISchoolService, mail adapters.IEmailAdapter, storage adapters.IFileUploader,
	validator common.IValidator, logger common.ILogger) SchoolController {
	validator.RegisterCustomValidator("validSchoolType", ValidateSchoolType)
	// instantiate subscription controller with dependencies
	controller := &SchoolController{
		service:         srv,
		mailProvider:    mail,
		serializer:      validator,
		logger:          logger,
		storageProvider: storage,
	}
	return *controller
}

func (cntrl *SchoolController) handleSchoolSetUp(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody SetUpSchoolDto

	if err := c.BodyParser(&requestBody); err != nil {
		log.Println(err.Error())
		return c.Status(422).JSON(common.APIResponse{Message: "invalid request body", Success: false, Detail: err.Error()})
	} else if !ValidateSchoolTypeByInput(requestBody.SchoolType) {
		return c.Status(422).JSON(common.APIResponse{Message: "invalid request body", Success: false, Detail: "School Type Not Valid"})
	}

	user, ok := c.Locals("user").(models.User)

	if !ok {
		// User not found in the context
		return fiber.ErrUnauthorized
	}

	// check if parent school was set up already
	_, err := cntrl.service.GetParentSchoolByOwnerID(ctx, user.ID)
	if err == nil {
		return c.Status(409).JSON(common.APIResponse{Message: "School Setup Failed", Success: false, Detail: "School Set Up already"})
	} else {
		log.Println(err.Error())
	}

	// upload logo and business registration documents
	logoUrl, err := storageService.UploadFile("logo", c)
	if err != nil {
		cntrl.logger.Error(ctx, "Error while setting up school",
			zap.String("additional_info", err.Error()))
	}
	documentUrl, err := storageService.UploadFile("document", c)
	if err != nil {
		cntrl.logger.Error(ctx, "Error while setting up school",
			zap.String("additional_info", err.Error()))
	}

	// set up new school
	school, err := cntrl.service.AddNewSchool(ctx, user.ID, requestBody.SchoolName, requestBody.Email, requestBody.SchoolType,
		requestBody.PhoneNumber, requestBody.Address, *requestBody.BrandColor, *logoUrl, *documentUrl)
	if err != nil && err == common.ErrDuplicate {
		return c.Status(409).JSON(common.APIResponse{Message: "School Setup Failed", Success: false, Detail: "School Set Up already"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error while setting up school",
			zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "School SetUp Failed", Success: false, Detail: "Internal Server Error"})
	}

	if err = cntrl.service.CreateDefaultSchoolRoles(ctx, school.ID); err != nil {
		cntrl.logger.Error(ctx, "Error while setting up school",
			zap.String("additional_info", err.Error()))
	}

	return c.Status(201).JSON(common.APIResponse{
		Message: "School SetUp Successfully",
		Success: true, Data: school})
}

func (cntrl *SchoolController) handleBranchAddition(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody AddSchoolBranchDto

	// validate request body
	if err := cntrl.serializer.ValidateRequestBody(c, &requestBody); err != nil {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: err.Error()})
	}

	user, ok := c.Locals("user").(models.User)

	if !ok {
		// User not found in the context
		return fiber.ErrUnauthorized
	}

	// get parent school created by super admin
	parentSchool, err := cntrl.service.GetParentSchoolByOwnerID(ctx, user.ID)
	if err != nil && err == common.ErrRecordNotFound {
		return c.Status(400).JSON(common.APIResponse{Message: "School Branch Addition Failed",
			Success: false, Detail: "School Has Not Yet Been Set Up"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error while adding school branch ", zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "School Branch Addition Failed", Success: false, Detail: "Internal Server Error"})
	}

	// add school branch
	branch, err := cntrl.service.AddSchoolBranch(ctx, parentSchool, requestBody.BranchName, requestBody.Email,
		requestBody.SchoolType, requestBody.Address, requestBody.PhoneNumber)
	if err != nil && err == common.ErrDuplicate {
		return c.Status(409).JSON(common.APIResponse{Message: "School Branch Addition Failed", Success: false, Detail: "School Branch Added already"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error while adding branch", zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "School Branch Addition Failed", Success: false, Detail: "Internal Server Error"})
	}

	if err = cntrl.service.CreateDefaultSchoolRoles(ctx, branch.ID); err != nil {
		cntrl.logger.Error(ctx, "Error while adding school branch",
			zap.String("additional_info", err.Error()))
	}

	return c.Status(201).JSON(common.APIResponse{
		Message: "School Branch Added",
		Success: true, Data: branch})
}

func (cntrl *SchoolController) handleSessionManagement(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var requestBody CreateSessionDto

	// validate request body
	if err := cntrl.serializer.ValidateRequestBody(c, &requestBody); err != nil {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: err.Error()})
	}

	user, ok := c.Locals("user").(models.User)

	if !ok {
		// User not found in the context
		return fiber.ErrUnauthorized
	}

	// get parent school created by super admin
	parentSchool, err := cntrl.service.GetParentSchoolByOwnerID(ctx, user.ID)
	if err != nil && err == common.ErrRecordNotFound {
		return c.Status(400).JSON(common.APIResponse{Message: "Session Creation Failed",
			Success: false, Detail: "School Has Not Yet Been Set Up"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error creating session ", zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "Session Creation Failed", Success: false, Detail: "Internal Server Error"})
	}

	// set up school session
	session, err := cntrl.service.CreateSession(ctx, parentSchool.ID, requestBody.SessionName, requestBody.SessionType,
		requestBody.SessionStart, requestBody.SessionEnd)
	if err != nil && err == common.ErrDuplicate {
		return c.Status(409).JSON(common.APIResponse{Message: "Session Creation Failed", Success: false, Detail: "School Branch Added already"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error creating session", zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "Session Creation Failed", Success: false, Detail: "Internal Server Error"})
	}

	return c.Status(201).JSON(common.APIResponse{
		Message: "Session Created",
		Success: true, Data: session})
}

func (cntrl *SchoolController) GetSchoolsOwned(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	user, ok := c.Locals("user").(models.User)

	if !ok {
		// User not found in the context
		return fiber.ErrUnauthorized
	}

	// get schools created by super admin
	schoolsOwned, err := cntrl.service.GetSchoolsByOwnerID(ctx, user.ID)
	if err != nil && err == common.ErrRecordNotFound {
		return c.Status(404).JSON(common.APIResponse{Message: "Failed To Fetch Schools",
			Success: false, Detail: "No School(s) set up by User"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error while fetching school", zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "Failed To Fetch Schools", Success: false, Detail: "Internal Server Error"})
	}

	return c.Status(200).JSON(common.APIResponse{
		Message: "Schools Fetched",
		Success: true, Data: schoolsOwned})
}

// get roles synonymous to job titles
func (cntrl *SchoolController) handleRolesFetching(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	// get roletype from query parameter
	roleType := c.Query("role_type", "Teaching Staff")
	schoolId := c.Query("school_id")

	if !ValidateRoleTypeByInput(roleType) {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: "role_type Not Valid"})
	} else if !common.ValidateUUIDByInput(schoolId) {
		return c.Status(422).JSON(common.APIResponse{Message: "Invalid Request Body", Success: false, Detail: "school_id Not Valid UUID Type"})
	}

	// get school roles
	roles, err := cntrl.service.GetSchoolRolesByType(ctx, schoolId, models.RoleType(roleType))
	if err != nil && err == common.ErrRecordNotFound {
		return c.Status(404).JSON(common.APIResponse{Message: "Failed To Get School Roles", Success: false, Detail: "School roles Not Yet Created"})
	} else if err != nil {
		cntrl.logger.Error(ctx, "Error while fetching school roles", zap.String("additional_info", err.Error()))
		return c.Status(500).JSON(common.APIResponse{Message: "Failed To Fetch School Roles", Success: false, Detail: "Internal Server Error"})
	}

	return c.Status(200).JSON(common.APIResponse{
		Message: "Fetched School Roles",
		Success: true,
		Data:    roles})
}
