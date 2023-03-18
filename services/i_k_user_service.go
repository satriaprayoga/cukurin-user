package services

import (
	"context"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/token"
)

type IKUserService interface {
	GetByEmailKUser(ctx context.Context, email string, usertype string) (result models.KUser, err error)
	ChangePassword(ctx context.Context, Payload token.Payload, DataChPwd models.ChangePassword) (err error)
	GetDataBy(ctx context.Context, Payload token.Payload, ID int) (result interface{}, err error)
	GetList(ctx context.Context, Payload token.Payload, queryparam models.ParamList) (result models.ResponseModelList, err error)
	Create(ctx context.Context, data *models.KUser) (err error)
}
