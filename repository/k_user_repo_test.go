package repo

import (
	"testing"

	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
)

func TestCreate(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()

}
