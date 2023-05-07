package session

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

type JWT struct {
	secret []byte
}

func (j *JWT) New(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedJWT, err := token.SignedString(j.secret)
	if err != nil {
		panic(err)
	}
	return signedJWT
}

func (j *JWT) Parse(tok string) (*Session, *jwt.Token, error) {
	token, pErr := jwt.Parse(tok, j.KeyFunc)
	if pErr != nil {
		return nil, nil, pErr
	}
	anyFirebaseToken, ftFound := token.Claims.(jwt.MapClaims)["firebaseToken"]
	if !ftFound {
		return nil, nil, fmt.Errorf("expecting firebaseToken")
	}
	firebaseToken, ok := anyFirebaseToken.(string)
	if !ok {
		return nil, nil, fmt.Errorf("expecting firebaseToken as string")
	}
	anyUserUUID, uFound := token.Claims.(jwt.MapClaims)["userUUID"]
	if !uFound {
		return nil, nil, fmt.Errorf("expecting userUUID")
	}
	stringUserUUID, ok := anyUserUUID.(string)
	if !ok {
		return nil, nil, fmt.Errorf("expecting userUUID as string")
	}
	userUUID, err := uuid.FromString(stringUserUUID)
	return NewSession(
		userUUID, firebaseToken,
	), token, err
}

func (j *JWT) KeyFunc(t *jwt.Token) (interface{}, error) { return j.secret, nil }

func New(secret []byte) *JWT {
	return &JWT{secret: secret}
}

const DefaultSecret = "TESTING"

func NewDefault() *JWT {
	return New([]byte(DefaultSecret))
}
