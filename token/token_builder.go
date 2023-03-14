package token

import "time"

type TokenBuilder interface {
	CreateToken(UserID string, Username string, UserType string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
