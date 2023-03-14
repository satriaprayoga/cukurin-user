package repo

import (
	"fmt"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/logging"
	"gorm.io/gorm"
)

type RepoKUser struct {
	Conn *gorm.DB
}

func NewRepoKUser(Conn *gorm.DB) *RepoKUser {
	return &RepoKUser{Conn}
}

func (db *RepoKUser) Create(data *models.KUser) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	q := db.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", q))
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *RepoKUser) Update(ID int, data interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)

	q := db.Conn.Model(models.KUser{}).Where("user_id=?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", q))
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}
