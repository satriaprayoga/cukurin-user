package services

import (
	"context"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/token"
)

type IAuthService interface {
	Logout(ctx context.Context, Payload token.Payload) error
	Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error)
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
}
