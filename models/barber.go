package models

import "time"

type Barber struct {
	BarberID       int       `json:"barber_id" gorm:"primary_key;auto_increment:true"`
	OwnerID        int       `json:"owner_id" gorm:"PRIMARY_KEY;type:integer"`
	BarberCd       string    `json:"barber_cd" gorm:"type:varchar(4)"`
	BarberName     string    `json:"barber_name" gorm:"type:varchar(60)"`
	Address        string    `json:"address" gorm:"type:varchar(150)"`
	FileID         int       `json:"file_id" gorm:"type:integer"`
	Latitude       float64   `json:"latitude" gorm:"type:float8"`
	Longitude      float64   `json:"longitude" gorm:"type:float8"`
	OperationStart time.Time `json:"operation_start" gorm:"type:timestamp(0) without time zone"`
	OperationEnd   time.Time `json:"operation_end" gorm:"type:timestamp(0) without time zone"`
	IsActive       bool      `json:"is_active" gorm:"type:boolean"`
	Model
}

type GeoBarber struct {
	Latitude  float64 `json:"latitude" gorm:"type:float8"`
	Longitude float64 `json:"longitude" gorm:"type:float8"`
}

type BarberPaket struct {
	BarberID  int       `json:"barber_id" gorm:"primary_key;type:integer"`
	PaketID   int       `json:"paket_id" gorm:"primary_key;type:integer"`
	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_Edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type BarberCapster struct {
	BarberID  int       `json:"barber_id" gorm:"primary_key;type:integer"`
	CapsterID int       `json:"capster_id" gorm:"primary_key;type:integer"`
	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_Edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type BarbersList struct {
	BarberID       int       `json:"barber_id"`
	BarberCd       string    `json:"barber_cd"`
	BarberName     string    `json:"barber_name"`
	Address        string    `json:"address"`
	OperationStart time.Time `json:"operation_start"`
	OperationEnd   time.Time `json:"operation_end"`
	IsActive       bool      `json:"is_active"`
	IsFavorit      bool      `json:"is_favorit"`
	Distance       float32   `json:"distance"`
	BarberRating   float32   `json:"barber_rating"`
	IsBarberOpen   bool      `json:"is_barber_open"`
	TotalUserOrder int       `json:"total_user_order"`
	//SaFileOutput
	GeoBarber
}
type BarbersUpdate struct {
	BarberName     string    `json:"barber_name" valid:"Required"`
	Address        string    `json:"address" valid:"Required"`
	FileID         int       `json:"file_id,omitempty"`
	OperationStart time.Time `json:"operation_start" valid:"Required"`
	OperationEnd   time.Time `json:"operation_end" valid:"Required"`
	IsActive       bool      `json:"is_active" valid:"Required"`
	GeoBarber
}

type BarbersPost struct {
	BarberName     string              `json:"barber_name" valid:"Required"`
	Address        string              `json:"address" valid:"Required"`
	FileID         int                 `json:"file_id,omitempty"`
	OperationStart time.Time           `json:"operation_start" valid:"Required"`
	OperationEnd   time.Time           `json:"operation_end" valid:"Required"`
	IsActive       bool                `json:"is_active" valid:"Required"`
	BarberPaket    []BarberPaketPost   `json:"barber_paket"`
	BarberCapster  []BarberCapsterPost `json:"barber_capster"`
	GeoBarber
}

type BarberPaketPost struct {
	PaketID int `json:"paket_id"`
}

type BarberCapsterPost struct {
	CapsterID int `json:"capster_id"`
}
