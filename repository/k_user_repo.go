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

func NewRepoKUser(Conn *gorm.DB) IKUserRepository {
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

func (db *RepoKUser) GetByAccount(account string, userType string) (result models.KUser, err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Where("email LIKE ? OR telp=? AND user_type=?", account, account, userType).Find(&result)
	logger.Query(fmt.Sprintf("%v", query))
	// logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	return result, err
}
