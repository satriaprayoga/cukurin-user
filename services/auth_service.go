package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/satriaprayoga/cukurin-user/models"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
	"github.com/satriaprayoga/cukurin-user/pkg/utils"
	repo "github.com/satriaprayoga/cukurin-user/repository"
	"github.com/satriaprayoga/cukurin-user/token"
)

type authService struct {
	repoKUser      repo.IKUserRepository
	repoKSession   repo.IKSessionRepository
	contextTimeOut time.Duration
}

func NewAuthService(a repo.IKUserRepository, b repo.IKSessionRepository, timeout time.Duration) IAuthService {
	return &authService{repoKUser: a, repoKSession: b, contextTimeOut: timeout}
}

func (a *authService) Logout(ctx context.Context, Payload token.Payload) error {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	err := a.repoKSession.DeleteByUserID(Payload.UserID)
	if err != nil {
		return err
	}
	return nil
}
func (a *authService) Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		User     models.KUser
		ksession models.KSession
	)

	CekData, err := a.repoKUser.GetByAccount(dataRegister.Account, dataRegister.UserType)
	if CekData.Email == dataRegister.Account {
		if CekData.IsActive {
			return output, errors.New("email sudah terdaftar")
		}
	}

	if dataRegister.Passwd != dataRegister.ConfirmPasswd {
		return output, errors.New("password dan confirm password harus sama")
	}

	User.Name = dataRegister.Name
	User.UserName = dataRegister.UserName
	User.JoinDate = time.Now()
	User.BirthOfDate = dataRegister.BirthOfDate
	User.UserType = dataRegister.UserType
	User.IsActive = false
	User.Email = dataRegister.Account
	User.Password, _ = utils.Hash(dataRegister.Passwd)

	if CekData.UserID > 0 && !CekData.IsActive {
		CekData.Name = User.Name
		CekData.Password = User.Password
		CekData.JoinDate = User.JoinDate
		CekData.UserType = User.UserType
		CekData.IsActive = User.IsActive
		CekData.Email = User.Email
		err = a.repoKUser.Update(CekData.UserID, CekData)
		if err != nil {
			return output, err
		}
	} else {
		User.UserEdit = dataRegister.UserName
		User.UserInput = dataRegister.UserName
		err = a.repoKUser.Create(&User)
		if err != nil {
			return output, err
		}
	}

	GenCode := utils.GenerateNumber(4)
	ksession.SessionID = GenCode
	ksession.UserID = User.UserID
	ksession.SessionType = "register"
	ksession.ExpiresAt = time.Now().Add(time.Hour * time.Duration(24))
	ksession.Account = User.Email
	err = a.repoKSession.Create(&ksession)
	if err != nil {
		return nil, err
	}
	out := map[string]interface{}{
		"otp":     GenCode,
		"account": User.Email,
	}

	return out, nil
}

func (a *authService) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		expireToken = settings.AppConfigSetting.JWTExpired
		ksession    models.KSession
	)

	DataUser, err := a.repoKUser.GetByAccount(dataLogin.Account, dataLogin.UserType)
	if err != nil {
		return nil, errors.New("email belum terdaftar")
	}

	if DataUser.UserType != "user" {
		return nil, errors.New("email belum terdaftar")
	}

	if !DataUser.IsActive {
		return nil, errors.New("akun belum aktif")
	}

	if !utils.ComparePassword(DataUser.Password, utils.GetPassword(dataLogin.Password)) {
		return nil, errors.New("password salah")
	}

	jwtToken, err := token.GenerateJwtToken(DataUser.UserID, DataUser.UserName, DataUser.UserType)
	if err != nil {
		return nil, err
	}
	ksession.SessionID = uuid.New().String()
	ksession.UserID = DataUser.UserID
	ksession.Account = dataLogin.Account
	ksession.ExpiresAt = time.Now().Add(time.Hour * time.Duration(expireToken))
	ksession.SessionType = "auth"

	err = a.repoKSession.Create(&ksession)
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
		"token":     jwtToken,
		"data_user": restUser,
	}

	return response, nil
}

func (a *authService) ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	if dataReset.Passwd != dataReset.ConfirmPasswd {
		return errors.New("password dan confirm Password harus sama")
	}

	DataUser, err := a.repoKUser.GetByAccount(dataReset.Account, "user")
	if err != nil {
		return err
	}

	if utils.ComparePassword(DataUser.Password, utils.GetPassword(dataReset.Passwd)) {
		return errors.New("password baru tidak boleh sama dengan yang lama")
	}

	DataUser.Password, _ = utils.Hash(dataReset.Passwd)

	err = a.repoKUser.UpdatePasswordByEmail(dataReset.Account, DataUser.Password)
	if err != nil {
		return err
	}

	return nil

}

func (a *authService) VerifyRegisterLogin(ctx context.Context, dataVerify *models.VerifyForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		expireToken = settings.AppConfigSetting.JWTExpired
		ksession    models.KSession
	)

	data, err := a.repoKSession.GetByAccount(dataVerify.Account)
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}
	if data.SessionID != dataVerify.VerifyCode {
		return nil, errors.New("otp yang anda masukkan salah")
	}

	DataUser, err := a.repoKUser.GetByAccount(dataVerify.Account, "user")
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}

	a.repoKSession.DeleteByUserID(DataUser.UserID)

	mUser := map[string]interface{}{
		"is_active": true,
	}

	err = a.repoKUser.Update(DataUser.UserID, mUser)
	if err != nil {
		return output, err
	}

	jwtToken, err := token.GenerateJwtToken(DataUser.UserID, DataUser.UserName, DataUser.UserType)
	if err != nil {
		return nil, err
	}
	ksession.SessionID = uuid.New().String()
	ksession.UserID = DataUser.UserID
	ksession.ExpiresAt = time.Now().Add(time.Hour * time.Duration(expireToken))
	ksession.SessionType = "auth"

	err = a.repoKSession.Create(&ksession)
	if err != nil {
		return nil, err
	}
	restUser := map[string]interface{}{
		"user_id":   DataUser.UserID,
		"email":     DataUser.Email,
		"telp":      DataUser.Telp,
		"user_name": DataUser.UserName,
		"user_type": DataUser.UserType,
		"name":      DataUser.Name,
	}
	response := map[string]interface{}{
		"token":     jwtToken,
		"data_user": restUser,
	}

	return response, nil
}

func (a *authService) VerifyRegister(ctx context.Context, dataVerify *models.VerifyForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	data, err := a.repoKSession.GetByAccount(dataVerify.Account)
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}
	if data.SessionID != dataVerify.VerifyCode {
		return nil, errors.New("otp yang anda masukkan salah")
	}

	DataUser, err := a.repoKUser.GetByAccount(dataVerify.Account, "user")
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}

	a.repoKSession.DeleteByUserID(DataUser.UserID)

	mUser := map[string]interface{}{
		"is_active": true,
	}

	err = a.repoKUser.Update(DataUser.UserID, mUser)
	if err != nil {
		return output, err
	}
	restUser := map[string]interface{}{
		"user_id":   DataUser.UserID,
		"email":     DataUser.Email,
		"telp":      DataUser.Telp,
		"user_name": DataUser.UserName,
		"user_type": DataUser.UserType,
		"name":      DataUser.Name,
	}

	return restUser, nil
}
