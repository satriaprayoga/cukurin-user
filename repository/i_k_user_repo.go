package repo

import "github.com/satriaprayoga/cukurin-user/models"

type IKUserRepository interface {
	Create(data *models.KUser) error
	Update(ID int, data interface{}) error
	GetByAccount(account string, userType string) (result models.KUser, err error)
}
