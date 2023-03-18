package repo

import "github.com/satriaprayoga/cukurin-user/models"

type IKUserRepository interface {
	Create(data *models.KUser) error
	Update(ID int, data interface{}) error
	GetByAccount(account string, userType string) (result models.KUser, err error)
	UpdatePasswordByEmail(Email string, Password string) error
	GetDataBy(ID int) (result *models.KUser, err error)
	GetList(queryparam models.ParamList) (result []*models.KUser, err error)
	Count(querparam models.ParamList) (result int, err error)
	Delete(ID int) error
}
