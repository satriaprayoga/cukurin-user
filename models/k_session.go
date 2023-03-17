package models

import (
	"time"

	"github.com/google/uuid"
)

type KSession struct {
	SessionID uuid.UUID `json:"session_id" gorm:"PRIMARY_KEY"`
	UserAgent string    `json:"user_agent" gorm:"type:varchar(30)"`
	ClientIp  string    `json:"client_ip" gorm:"type:varchar(32)"`
	ExpiresAt time.Time `json:"join_date" gorm:"type:timestamp(0)"`
	UserID    int       `json:"user_id" gorm:"type:integer;not null"`
	Model
}
