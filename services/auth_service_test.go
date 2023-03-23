package services

import (
	"context"
	"testing"
	"time"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/logging"
	"github.com/satriaprayoga/cukurin-user/pkg/sessions"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	repo "github.com/satriaprayoga/cukurin-user/repository"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {

	settings.Setup("../config/config.json")
	database.Setup()
	sessions.Setup()
	logging.Setup()

	var (
		timeOut      = settings.AppConfigSetting.Server.ReadTimeOut
		ctx          = context.Background()
		registerForm models.RegisterForm
	)
	repoKUser := repo.NewRepoKUser(database.Conn)
	authService := NewAuthService(repoKUser, time.Second*time.Duration(timeOut))
	registerForm.Account = "satria.prayoga@gmail.com"
	registerForm.BirthOfDate, _ = utils.GetDayOfBirth(1987, 05, 04, "2022-03-12")
	registerForm.Name = "Gilang Satria"
	registerForm.UserName = "gsprayoga"
	registerForm.Passwd = "asdqwe123"
	registerForm.ConfirmPasswd = "asdqwe123"
	registerForm.UserType = "user"
	data, err := authService.Register(ctx, registerForm)
	require.NoError(t, err)
	require.NotNil(t, data)
}
