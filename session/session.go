package session

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

const (
	DefaultIssuer          = "MESSAGE-API" // FIXME:
	DefaultSubject         = "AUTH_TOKEN"  //FIXME:
	DefaultAudience        = "APP_USERS"   //FIXME:
	DefaultExpirationDelta = 30 * 24 * time.Hour
)

type Session struct {
	UserUUID      uuid.UUID `json:"userUUID"`
	FirebaseToken string    `json:"firebaseToken"`
	issuedAt      time.Time
	expiration    time.Time
	notBefore     time.Time
}

func (s *Session) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(s.expiration), nil
}

func (s *Session) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(s.issuedAt), nil
}

func (s *Session) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(s.notBefore), nil
}

func (s *Session) GetIssuer() (string, error) {
	return DefaultIssuer, nil
}

func (s *Session) GetSubject() (string, error) {
	return DefaultSubject, nil
}

func (s *Session) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{DefaultAudience}, nil
}

func NewSession(userUUID uuid.UUID, firebaseToken string) *Session {
	now := time.Now()
	return &Session{
		UserUUID:      userUUID,
		FirebaseToken: firebaseToken,
		issuedAt:      now,
		expiration:    now.Add(DefaultExpirationDelta),
		notBefore:     now,
	}
}
