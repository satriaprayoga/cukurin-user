package repo

import (
	"fmt"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/logging"
	"gorm.io/gorm"
)

type repoBarber struct {
	Conn *gorm.DB
}

func NewRepoBarber(Conn *gorm.DB) IBarberRepository {
	return &repoBarber{Conn}
}

func (b *repoBarber) GetDataBy(ID int) (result *models.Barber, err error) {
	var (
		logger = logging.Logger{}
		barber = &models.Barber{}
	)
	query := b.Conn.Where("barber_id=?", ID).Find(barber)
	logger.Query(fmt.Sprintf("%v", query.Statement.SQL.String()))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return barber, nil
}
