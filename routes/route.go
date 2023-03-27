package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-user/controllers"
	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	repo "github.com/satriaprayoga/cukurin-user/repository"
	"github.com/satriaprayoga/cukurin-user/services"
)

type AppRoutes struct {
	E *echo.Echo
}

func (a *AppRoutes) Setup() {
	timeoutContext := time.Duration(settings.AppConfigSetting.Server.ReadTimeOut) * time.Second

	repoKUser := repo.NewRepoKUser(database.Conn)
	repoFileUpload := repo.NewRepoFileUpload(database.Conn)

	authService := services.NewAuthService(repoKUser, timeoutContext)
	controllers.NewAuthController(a.E, authService)

	kUserService := services.NewKUserService(repoKUser, timeoutContext)
	controllers.NewKUserController(a.E, kUserService)

	fileUploadService := services.NewFileUploadSevice(repoFileUpload, timeoutContext)
	controllers.NewFileUploadController(a.E, fileUploadService)
}
