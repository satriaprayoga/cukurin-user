package token

import (
	"context"

	"github.com/satriaprayoga/cukurin-user/models"
)

type TokenService interface {
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
}
