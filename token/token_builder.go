package token

type TokenBuilder interface {
	CreateToken(UserID string, Username string, UserType string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
