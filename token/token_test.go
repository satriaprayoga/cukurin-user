package token

import (
	"testing"
	"time"

	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
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
