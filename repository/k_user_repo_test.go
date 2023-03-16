package repo

import (
	"testing"

	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func TestCreate(t *testing.T) {
	settings.Setup("../config/config.json")
	database.Setup()
}
