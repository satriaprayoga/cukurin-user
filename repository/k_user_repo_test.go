package repo

import (
	"testing"
	"time"

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

	var (
		now time.Time
	)
	pwd, _ := utils.Hash("t3stPassword")
	now = time.Now()
	kUser := &models.KUser{
		Name:     utils.RandomName(),
		UserName: utils.RandomUserName(),
		Email:    utils.RandomEmail(),
		Telp:     utils.RandomPhone(),
		Password: pwd,
		JoinDate: now,
		IsActive: false,
		UserType: "user",
	}

	repoKUser := NewRepoKUser(database.Conn)
	err := repoKUser.Create(kUser)
	require.NoError(t, err)

}

func TestGetAccount(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
	logging.Setup()

	repoKUser := NewRepoKUser(database.Conn)
	data, err := repoKUser.GetByAccount("cuaxhx@mail.com", "user")
	require.NoError(t, err)
	require.NotNil(t, data)
	require.Equal(t, "cuaxhx@mail.com", data.Email)

}
