package token

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

func TestJWTBuilder(t *testing.T) {
	settings.Setup("../config/config.json")
	_, err := NewPayload(utils.GenerateString(5), utils.RandomUserName(), "user")
	require.NoError(t, err)
	issuedAt := time.Now()
	token_builder := NewJWTBuilder(settings.AppConfigSetting.App.JwtSecret)
	tkn, err := token_builder.CreateToken(utils.GenerateString(5), utils.RandomUserName(), "user")
	require.NoError(t, err)
	require.NotEmpty(t, tkn)

	payload, err := token_builder.VerifyToken(tkn)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.WithinDuration(t, issuedAt, payload.ExpiresAt, time.Duration(settings.AppConfigSetting.JWTExpired)*time.Hour)

}

func TestLogin(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()

	var (
		ctx       = context.Background()
		dataLogin = &models.LoginForm{
			Account:  "hthctc@mail.com",
			Password: "t3stPassword",
			UserType: "user",
		}
	)
	repoKUser := repo.NewRepoKUser(database.Conn)
	tokenService := NewJwtTokenService(repoKUser, time.Duration(2)*time.Second)
	resp, err := tokenService.Login(ctx, dataLogin)

	require.NoError(t, err)
	require.NotNil(t, resp)

}
