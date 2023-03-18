package usecase

import (
	"context"
	"time"

	"github.com/satriaprayoga/cukurin-user/models"
	repo "github.com/satriaprayoga/cukurin-user/repository"
)

type KUserService struct {
	kuserrepo      repo.IKUserRepository
	contextTimeOut time.Duration
}

func NewKUserService(kUserRepo repo.IKUserRepository, cto time.Duration) *KUserService {
	return &KUserService{kuserrepo: kUserRepo, contextTimeOut: cto}
}

func (r *KUserService) GetByEmailKUser(ctx context.Context, email string, usertype string) (result models.KUser, err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	kuser := models.KUser{}
	result, err = r.kuserrepo.GetByAccount(email, usertype)
	if err != nil {
		return kuser, err
	}
	return result, nil

}

func (r *KUserService) Create(ctx context.Context, data *models.KUser) (err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	err = r.kuserrepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}
