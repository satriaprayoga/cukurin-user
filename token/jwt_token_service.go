package token

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	repo "github.com/satriaprayoga/cukurin-user/repository"
)

type JwtTokenService struct {
	kuserrepo      repo.IKUserRepository
	tokenBuilder   TokenBuilder
	contextTimeOut time.Duration
}

func NewJwtTokenService(k repo.IKUserRepository, cto time.Duration) TokenService {
	var secret = settings.AppConfigSetting.App.JwtSecret
	return &JwtTokenService{
		kuserrepo:      k,
		contextTimeOut: cto,
		tokenBuilder:   NewJWTBuilder(secret),
	}
}

func (j *JwtTokenService) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	//var (
	//	logger = logging.Logger{}
	//)
	_, cancel := context.WithTimeout(ctx, j.contextTimeOut)
	defer cancel()

	DataUser, err := j.kuserrepo.GetByAccount(dataLogin.Account, dataLogin.UserType)
	if err != nil {
		return nil, errors.New("email anda belum terdaftar")
	}
	// if !DataUser.IsActive {
	// 	return nil, errors.New("Account anda belum aktif. Silahkan Register ulang dengan email yang sama.")
	// }
	if !utils.ComparePassword(DataUser.Password, utils.GetPassword(dataLogin.Password)) {
		return nil, errors.New("password yang anda masukkan salah, silahkan coba lagi")
	}
	t, err := j.tokenBuilder.CreateToken(strconv.Itoa(DataUser.UserID), DataUser.UserName, DataUser.UserType)
	if err != nil {
		return nil, err
	}
	restUser := map[string]interface{}{
		"user_id":   DataUser.UserID,
		"email":     DataUser.Email,
		"telp":      DataUser.Telp,
		"user_name": DataUser.Name,
		"user_type": DataUser.UserType,
	}

	response := map[string]interface{}{
		"token":     t,
		"data_user": restUser,
	}

	//jsonResp, err := json.Marshal(response)
	//if err != nil {
	//	logger.Error("cannot convert to json", response)
	//}
	//logger.Info("token response: ", string(jsonResp))
	return response, nil
}
