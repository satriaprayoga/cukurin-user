package token

import (
	"testing"

	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreatePayload(t *testing.T) {
	settings.Setup()
	_, err := NewPayload(utils.GenerateString(5), utils.RandomUserName(), "user")
	require.NoError(t, err)
}
