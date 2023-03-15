package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id,omitempty"`
	Username  string    `json:"user_name,omitempty"`
	UserType  string    `json:"user_type,omitempty"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
	Issuer    string    `json:"iss"`
}

func NewPayload(UserID string, Username string, UserType string) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	issuer := settings.AppConfigSetting.App.Issuer
	expiredTime := settings.AppConfigSetting.JWTExpired
	payload := &Payload{
		ID:        tokenId,
		Username:  Username,
		UserID:    UserID,
		UserType:  UserType,
		Issuer:    issuer,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiredTime)),
	}

	return payload, nil
}

func (P *Payload) Valid() error {
	if time.Now().After(P.ExpiresAt) {
		return errors.New("expired token")
	}
	return nil
}
