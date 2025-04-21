package main

import (
	ctx "context"
	"flag"
	"fmt"
	"os"
	"sample-scm-backend/api/routes"
	"sample-scm-backend/config"
	"sample-scm-backend/services"
	"time"

	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/distribution"
	"github.com/AMETORY/ametory-erp-modules/file"
	"github.com/AMETORY/ametory-erp-modules/inventory"
	"github.com/AMETORY/ametory-erp-modules/order"
	"github.com/AMETORY/ametory-erp-modules/shared/audit_trail"
	"github.com/AMETORY/ametory-erp-modules/shared/indonesia_regional"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := ctx.Background()
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db, err := services.InitDB(cfg)
	if err != nil {
		panic(err)
	}
	redisClient := services.InitRedis()
	websocket := services.InitWS()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3036",
		},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Origin",
			"API-KEY",
			"Currency-Code",
			"Cache-Control",
			"X-Requested-With",
			"Content-Disposition",
			"Content-Description",
			"ID-Company",
			"start-date",
			"end-date",
			"ID-Distributor",
			"timezone",
		},
		ExposeHeaders: []string{"Content-Length", "Content-Disposition", "Content-Description"},
	}))

	skipMigration := true

	if os.Getenv("MIGRATION") != "" {
		skipMigration = false
	}

	erpContext := context.NewERPContext(db, nil, &ctx, skipMigration)
	appService := services.NewAppService(erpContext, cfg, redisClient, websocket)
	erpContext.AppService = appService

	authService := auth.NewAuthService(erpContext)
	erpContext.AuthService = authService

	fileService := file.NewFileService(erpContext, cfg.Server.BaseURL)
	erpContext.FileService = fileService

	rbacSrv := auth.NewRBACService(erpContext)
	erpContext.RBACService = rbacSrv

	inventorySrv := inventory.NewInventoryService(erpContext)
	erpContext.InventoryService = inventorySrv

	auditTrailSrv := audit_trail.NewAuditTrailService(erpContext)

	orderService := order.NewOrderService(erpContext)
	erpContext.OrderService = orderService

	distributionSrv := distribution.NewDistributionService(erpContext, auditTrailSrv, inventorySrv, orderService)
	erpContext.DistributionService = distributionSrv

	indonesiaRegSrv := indonesia_regional.NewIndonesiaRegService(erpContext)
	erpContext.IndonesiaRegService = indonesiaRegSrv

	googleSrv := google.NewGoogleAPIService(erpContext, cfg.Google.APIKey)

	erpContext.AddThirdPartyService("google", googleSrv)

	flagCreateUser := flag.Bool("create-user", false, "create user with default permissions")
	flagEmail := flag.String("email", "", "user email")
	flagFullName := flag.String("full-name", "", "user full name")
	flagPassword := flag.String("password", "", "user password")
	flag.Parse()

	if *flagCreateUser {
		if *flagEmail == "" {
			fmt.Println("email is required")
			return
		}
		if *flagPassword == "" {
			fmt.Println("password is required")
			return
		}
		if *flagFullName == "" {
			fmt.Println("full name is required")
			return
		}
		hashedPassword, err := models.HashPassword(*flagPassword)
		if err != nil {
			fmt.Println(err)
			return
		}
		now := time.Now()
		username := utils.CreateUsernameFromFullName(*flagFullName)
		user := &models.UserModel{
			Email:      *flagEmail,
			Password:   hashedPassword,
			Username:   username,
			VerifiedAt: &now,
			FullName:   *flagFullName,
		}
		err = appService.CreateUserWithDefaultPermissions(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("user created")
		return
	}

	v1 := r.Group("/api/v1")

	r.Static("/assets/files", "./assets/files")
	routes.SetupWSRoutes(v1, erpContext)
	routes.SetupAuthRoutes(v1, erpContext)
	routes.SetupProductRoutes(v1, erpContext)
	routes.SetupProductCategoryRoutes(v1, erpContext)
	routes.SetupStockMovementRoutes(v1, erpContext)
	routes.SetupUnitRoutes(v1, erpContext)
	routes.SetupStorageRoutes(v1, erpContext)
	routes.SetupStockOpnameRoutes(v1, erpContext)
	routes.SetupLogisticRoutes(v1, erpContext)
	routes.SetupCommonRoutes(v1, erpContext)
	routes.SetupRegionalRoutes(v1, erpContext)

	if os.Getenv("GEN_PERMISSIONS") != "" {
		appService.GenerateDefaultPermissions()
	}

	r.Run(":" + config.App.Server.Port)
}
