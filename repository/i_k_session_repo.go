package repo

import (
	"github.com/satriaprayoga/cukurin-user/models"
)

type IKSessionRepository interface {
	Create(data *models.KSession) error
	GetByUserID(UserID int) (output *models.KSession, err error)
	DeleteByUserID(UserID int) (err error)
}
