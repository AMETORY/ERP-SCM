package handlers

import (
	"errors"
	"sample-scm-backend/services"

	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/file"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommonHandler struct {
	ctx         *context.ERPContext
	appService  *services.AppService
	rbacService *auth.RBACService
	authService *auth.AuthService
	fileService *file.FileService
}

func NewCommonHandler(ctx *context.ERPContext) *CommonHandler {

	appService, ok := ctx.AppService.(*services.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	rbacService, ok := ctx.RBACService.(*auth.RBACService)
	if !ok {
		panic("RBACService is not instance of auth.RBACService")
	}
	authService, ok := ctx.AuthService.(*auth.AuthService)
	if !ok {
		panic("AuthService is not instance of auth.AuthService")
	}
	fileService, ok := ctx.FileService.(*file.FileService)
	if !ok {
		panic("FileService is not instance of file.FileService")
	}

	return &CommonHandler{
		ctx:         ctx,
		appService:  appService,
		rbacService: rbacService,
		authService: authService,
		fileService: fileService,
	}
}

func (h *CommonHandler) GetRolesHandler(c *gin.Context) {
	roles, err := h.rbacService.GetAllRoles(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	items := roles.Items.(*[]models.RoleModel)
	newItems := make([]models.RoleModel, 0)
	for _, v := range *items {
		if !v.IsSuperAdmin {
			v.Permissions = nil
			newItems = append(newItems, v)
		}
	}
	roles.Items = &newItems
	c.JSON(200, gin.H{"data": roles})
}

func (h *CommonHandler) UploadFileHandler(c *gin.Context) {
	h.ctx.Request = c.Request

	fileObject := models.FileModel{}
	refID, _ := c.GetPostForm("ref_id")
	refType, _ := c.GetPostForm("ref_type")
	skipSave := false
	skipSaveStr, _ := c.GetPostForm("skip_save")
	if skipSaveStr == "true" || skipSaveStr == "1" {
		skipSave = true
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fileByte, err := utils.FileHeaderToBytes(file)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	filename := file.Filename

	fileObject.FileName = utils.FilenameTrimSpace(filename)
	fileObject.RefID = refID
	fileObject.RefType = refType
	fileObject.SkipSave = skipSave

	if err := h.fileService.UploadFile(fileByte, "local", "files", &fileObject); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "File uploaded successfully", "data": fileObject})
}

func checkSuperAdmin(erpContext *context.ERPContext, userID string, companyID string) (bool, error) {
	var admin models.UserModel

	// Cari pengguna beserta peran dan izin
	if err := erpContext.DB.Preload("Roles", func(db *gorm.DB) *gorm.DB {
		return db.Where("company_id = ?", companyID)
	}).First(&admin, "id = ?", userID).Error; err != nil {
		return false, errors.New("admin not found")
	}

	// Periksa izin
	for _, role := range admin.Roles {
		if role.IsSuperAdmin {
			return true, nil
		}
	}

	return false, nil
}
