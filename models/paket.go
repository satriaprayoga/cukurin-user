package models

import "time"

type Paket struct {
	PaketID     int        `json:"paket_id" gorm:"primary_key;auto_increment:true"`
	OwnerID     int        `json:"barber_id" gorm:"PRIMARY_KEY;type:integer"`
	PaketName   string     `json:"paket_name" gorm:"type:varchar(60)"`
	Desc        string     `json:"descs" gorm:"type:varchar(200)"`
	DurasiStart int        `json:"durasi_start" gorm:"type:integer"`
	DurasiEnd   int        `json:"durasi_end" gorm:"type:integer"`
	Price       float32    `json:"price" gorm:"type:numeric(20,4)"`
	IsActive    bool       `json:"is_active" gorm:"type:boolean"`
	IsPromo     bool       `json:"is_promo" gorm:"type:boolean"`
	PromoPrice  float32    `json:"promo_price" gorm:"type:numeric(20,4)"`
	PromoStart  *time.Time `json:"promo_start" gorm:"type:timestamp(0) without time zone"`
	PromoEnd    *time.Time `json:"promo_end" gorm:"type:timestamp(0) without time zone"`
	Model
}

type DataPaket struct {
	// OwnerID     int       `json:"owner_id" valid:"Required"`
	PaketName   string    `json:"paket_name" valid:"Required"`
	Descs       string    `json:"descs,omitempty"`
	DurasiStart int       `json:"durasi_start,omitempty"`
	DurasiEnd   int       `json:"durasi_end,omitempty"`
	Price       float32   `json:"price" valid:"Required"`
	IsActive    bool      `json:"is_active"`
	IsPromo     bool      `json:"is_promo,omitempty"`
	PromoPrice  float32   `json:"promo_price,omitempty"`
	PromoStart  time.Time `json:"promo_start,omitempty"`
	PromoEnd    time.Time `json:"promo_end,omitempty"`
}
