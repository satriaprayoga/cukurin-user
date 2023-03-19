package models

import "time"

type LoginForm struct {
	Account  string `json:"account" valid:"Required"`
	Password string `json:"pwd" valid:"Required"`
	UserType string `json:"user_type" valid:"Required"`
	//FcmToken string `json:"fcm_token" valid:"Required"`
}

type RegisterForm struct {
	Name          string    `json:"name" valid:"Required"`
	UserName      string    `json:"user_name" valid:"Required"`
	BirthOfDate   time.Time `json:"birth_of_date"`
	Account       string    `json:"account" valid:"Required"`
	Passwd        string    `json:"pwd" valid:"Required"`
	ConfirmPasswd string    `json:"confirm_pwd" valid:"Required"`
	UserType      string    `json:"user_type"`
}
