package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/satriaprayoga/cukurin-user/pkg/settings"
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	UserID   string    `json:"user_id,omitempty"`
	Username string    `json:"user_name,omitempty"`
	UserType string    `json:"user_type,omitempty"`
	jwt.StandardClaims
}

func NewPayload(UserID string, Username string, UserType string) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	issuer := settings.AppConfigSetting.App.Issuer
	expiredTime := settings.AppConfigSetting.JWTExpired
	payload := &Payload{
		ID:       tokenId,
		Username: Username,
		UserID:   UserID,
		UserType: UserType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix(),
		},
	}

	return payload, nil
}
