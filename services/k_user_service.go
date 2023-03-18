package services

import (
	"context"
	"errors"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/database"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	repo "github.com/satriaprayoga/cukurin-user/repository"
	"github.com/satriaprayoga/cukurin-user/token"
)

type KUserService struct {
	kuserrepo      repo.IKUserRepository
	contextTimeOut time.Duration
}

func NewKUserService(kUserRepo repo.IKUserRepository, cto time.Duration) IKUserService {
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

func (r *KUserService) ChangePassword(ctx context.Context, Payload token.Payload, DataChPwd models.ChangePassword) (err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	ID, _ := strconv.Atoi(Payload.UserID)
	dataUser, err := r.kuserrepo.GetDataBy(ID)
	if err != nil {
		return err
	}

	if !utils.ComparePassword(dataUser.Password, utils.GetPassword(DataChPwd.OldPassword)) {
		return errors.New("password lama anda salah")
	}

	if DataChPwd.NewPassword != DataChPwd.ConfirmPassword {
		return errors.New("password dan confirm password tidaks sama")
	}

	if utils.ComparePassword(dataUser.Password, utils.GetPassword(DataChPwd.NewPassword)) {
		return errors.New("password baru tidak boleh sama dengan yang lama")
	}

	DataChPwd.NewPassword, _ = utils.Hash(DataChPwd.NewPassword)

	err = r.kuserrepo.UpdatePasswordByEmail(dataUser.Email, DataChPwd.NewPassword)
	if err != nil {
		return err
	}
	return nil
}

func (r *KUserService) GetDataBy(ctx context.Context, Payload token.Payload, ID int) (result interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	DataUser, err := r.kuserrepo.GetDataBy(ID)
	if err != nil {
		if err != models.ErrNotFound {
			return result, err
		}
	}

	response := map[string]interface{}{
		"user_id":       DataUser.UserID,
		"user_name":     DataUser.UserName,
		"birth_of_date": DataUser.BirthOfDate,
		"email":         DataUser.Email,
		"telp":          DataUser.Telp,
	}
	return response, nil
}

func (r *KUserService) GetList(ctx context.Context, Payload token.Payload, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	var kUser = models.KUser{}
	if queryparam.Search != "" {
		value := reflect.ValueOf(kUser)
		types := reflect.TypeOf(&kUser)
		queryparam.Search = database.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = r.kuserrepo.GetList(queryparam)
	if err != nil {
		return result, err
	}
	result.Total, err = r.kuserrepo.Count(queryparam)
	if err != nil {
		return result, err
	}
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

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
