package models

type LoginForm struct {
	Account  string `json:"account" valid:"Required"`
	Password string `json:"pwd" valid:"Required"`
	UserType string `json:"user_type" valid:"Required"`
	//FcmToken string `json:"fcm_token" valid:"Required"`
}
