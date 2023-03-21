package services

import (
	"context"
	"testing"
	"time"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/logging"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	repo "github.com/satriaprayoga/cukurin-user/repository"
	"github.com/stretchr/testify/require"
)

func TestAuthRegister(t *testing.T) {
	var (
		ctx          = context.Background()
		dataRegister models.RegisterForm
	)
	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()
	repoKUser := repo.NewRepoKUser(database.Conn)
	repoKSession := repo.NewRepoKSession(database.Conn)
	//var expireToken = settings.AppConfigSetting.JWTExpired
	authService := NewAuthService(repoKUser, repoKSession, time.Duration(time.Duration(3)*time.Millisecond))

	dataRegister.Name = "Gilang SP"
	dataRegister.UserName = "gsprayoga"
	dataRegister.Account = "satria.prayoga@gmail.com"
	dataRegister.Passwd = "asdqwe123"
	dataRegister.ConfirmPasswd = "asdqwe123"
	dataRegister.UserType = "user"
	dob, _ := utils.GetDayOfBirth(1987, 05, 04, "2020-10-05")
	dataRegister.BirthOfDate = dob
	out, err := authService.Register(ctx, dataRegister)
	require.NoError(t, err)
	require.NotNil(t, out)
}

func TestVerifyRegisterLogin(t *testing.T) {
	var (
		ctx        = context.Background()
		dataVerify *models.VerifyForm
	)
	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()
	repoKUser := repo.NewRepoKUser(database.Conn)
	repoKSession := repo.NewRepoKSession(database.Conn)
	//var expireToken = settings.AppConfigSetting.JWTExpired
	authService := NewAuthService(repoKUser, repoKSession, time.Duration(time.Duration(3)*time.Millisecond))
	dataVerify = &models.VerifyForm{
		Account:    "satria.prayoga@gmail.com",
		VerifyCode: "4470",
	}
	out, err := authService.VerifyRegisterLogin(ctx, dataVerify)
	require.NoError(t, err)
	require.NotNil(t, out)

}

func TestVerifyRegister(t *testing.T) {
	var (
		ctx        = context.Background()
		dataVerify *models.VerifyForm
	)
	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()
	repoKUser := repo.NewRepoKUser(database.Conn)
	repoKSession := repo.NewRepoKSession(database.Conn)
	//var expireToken = settings.AppConfigSetting.JWTExpired
	authService := NewAuthService(repoKUser, repoKSession, time.Duration(time.Duration(3)*time.Millisecond))
	dataVerify = &models.VerifyForm{
		Account:    "satria.prayoga@gmail.com",
		VerifyCode: "2545",
	}
	out, err := authService.VerifyRegister(ctx, dataVerify)
	require.NoError(t, err)
	require.NotNil(t, out)

}

func TestAuthLogin(t *testing.T) {
	var (
		ctx = context.Background()
	)

	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()
	repoKUser := repo.NewRepoKUser(database.Conn)
	repoKSession := repo.NewRepoKSession(database.Conn)
	//var expireToken = settings.AppConfigSetting.JWTExpired
	authService := NewAuthService(repoKUser, repoKSession, time.Duration(time.Duration(3)*time.Millisecond))
	dataLogin := &models.LoginForm{
		Account:  "satria.prayoga@gmail.com",
		Password: "asdqwe123",
		UserType: "user",
	}

	out, err := authService.Login(ctx, dataLogin)
	require.NoError(t, err)
	require.NotNil(t, out)

}
