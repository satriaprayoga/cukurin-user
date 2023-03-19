package token

import (
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
)

func GenerateJwtToken(UserID int, Username string, UserType string) (string, error) {
	var (
		tokenBuilder TokenBuilder
		secret       string
	)

	secret = settings.AppConfigSetting.App.JwtSecret
	tokenBuilder = NewJWTBuilder(secret)
	t, err := tokenBuilder.CreateToken(UserID, Username, UserType)
	return t, err

}
