package services

import (
	"testing"

	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
)

func TestAuthRegister(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
	//repoKUser := repo.NewRepoKUser(database.Conn)
	//repoKSession := repo.NewRepoKSession(database.Conn)
	//var expireToken = settings.AppConfigSetting.JWTExpired
	//authService := NewAuthService(repoKUser, repoKSession, time.Duration(time.Duration(3)*time.Millisecond))
	//authService.Register()
}
