package repo

import "github.com/satriaprayoga/cukurin-user/models"

type IBarberRepository interface {
	GetDataBy(ID int) (result *models.Barber, err error)
}
