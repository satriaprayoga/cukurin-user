package services

import (
	"context"
	"time"

	repo "github.com/satriaprayoga/cukurin-user/repository"
	"github.com/satriaprayoga/cukurin-user/token"
)

type authService struct {
	repoKUser      repo.IKUserRepository
	repoKSession   repo.IKSessionRepository
	contextTimeOut time.Duration
}

func NewAuthService(a repo.IKUserRepository, b repo.IKSessionRepository, timeout time.Duration) *authService {
	return &authService{repoKUser: a, repoKSession: b, contextTimeOut: timeout}
}

func (a *authService) Logout(ctx context.Context, Payload token.Payload) error {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	err := a.repoKSession.DeleteByUserID(Payload.UserID)
	if err != nil {
		return err
	}
	return nil
}
