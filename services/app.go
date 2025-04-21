package services

import (
	"errors"
	"sample-scm-backend/config"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/go-redis/redis/v8"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type AppService struct {
	ctx       *context.ERPContext
	Config    *config.Config
	Redis     *redis.Client
	Websocket *melody.Melody
}

func NewAppService(ctx *context.ERPContext, config *config.Config, redis *redis.Client, ws *melody.Melody) *AppService {
	if !ctx.SkipMigration {
		ctx.DB.AutoMigrate(
		// &app_models.AppModel{},
		// &app_models.CustomSettingModel{},
		)
	}
	return &AppService{
		ctx:       ctx,
		Config:    config,
		Redis:     redis,
		Websocket: ws,
	}
}

func (a AppService) GenerateDefaultPermissions() []models.PermissionModel {
	var (
		cruds    = []string{"create", "read", "update", "delete"}
		services = map[string][]map[string][]string{
			"auth": {{"user": cruds, "admin": cruds, "rbac": cruds}},
			"inventory": {
				{"purchase": cruds},
				{"purchase_return": cruds},
				{"product": cruds},
				{"product_category": cruds},
				{"price_category": cruds},
				{"product_attribute": cruds},
				{"warehouse": cruds},
				{"unit": cruds},
				{"stock_movement": cruds},
				{"stock_opname": cruds},
			},
			"distribution": {
				{"logistic": []string{
					"read-distribution-event",
					"list-distribution-event",
					"create-distribution-event",
					"create-shipment",
					"ready-to-ship",
					"process-shipment",
					"create-shipment-leg",
					"start-shipment-leg",
					"arrived-shipment-leg",
					"add-tracking-event",
					"generate-shipment-report",
					"generate-distributor-event-report",
					"report-lost-damage",
				}},
				{"storage": []string{
					"list-warehouse",
					"create-warehouse",
					"delete-warehouse",
					"update-warehouse",
					"create-location",
					"update-location",
					"delete-location",
					"get-location-detail",
					"get-locations",
				}},
			},
		}
	)

	return a.generatePermissions(services)
}

func (a AppService) generatePermissions(services map[string][]map[string][]string) []models.PermissionModel {

	// create Default SuperAdmin Role
	var role models.RoleModel
	err := a.ctx.DB.First(&role, "name = ?", "Super Admin").Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		role = models.RoleModel{
			Name:         "Super Admin",
			IsSuperAdmin: true,
		}
		a.ctx.DB.Create(&role)
	}

	var permissions []models.PermissionModel

	for service, modules := range services {
		for _, module := range modules {
			for key, actions := range module {
				for _, action := range actions {
					var permission models.PermissionModel
					err := a.ctx.DB.First(&permission, "name = ?", service+":"+key+":"+action).Error
					if errors.Is(err, gorm.ErrRecordNotFound) {
						permission.Name = service + ":" + key + ":" + action
						a.ctx.DB.Create(&permission)
					}
					permissions = append(permissions, permission)
				}
			}
		}
	}
	return permissions
}

func (a AppService) CreateUserWithDefaultPermissions(user *models.UserModel) error {
	role := models.RoleModel{}
	err := a.ctx.DB.First(&role, "name = ?", "Super Admin").Error
	if err != nil {
		return err
	}
	if err := a.ctx.DB.Create(user).Error; err != nil {
		return err
	}
	return a.ctx.DB.Model(user).Association("Roles").Append(&role)
}
