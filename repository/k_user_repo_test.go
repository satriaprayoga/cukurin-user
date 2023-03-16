package repo

import (
	"testing"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/logging"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()

	data := &models.KUser{
		UserName: utils.RandomUserName(),
		Name:     utils.RandomName(),
		Email:    utils.RandomEmail(),
		Telp:     utils.RandomPhone(),
		IsActive: false,
		UserType: "user",
		Password: utils.RandomPassword(),
	}
	repoKUser := NewRepoKUser(database.Conn)
	err := repoKUser.Create(data)
	require.NoError(t, err)

}

func TestGetAccount(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()

	repoKUser := NewRepoKUser(database.Conn)
	data, err := repoKUser.GetByAccount("hthctc@mail.com", "user")
	require.NoError(t, err)
	require.NotNil(t, data)
	require.Equal(t, "hthctc@mail.com", data.Email)
}
