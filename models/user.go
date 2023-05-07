package models

import (
	"fmt"
	"net/mail"
	"regexp"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/common/random"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	DefaultPasswordCost    = 12
	DefaultExpirationDelta = 30 * 24 * time.Hour
)

type User struct {
	Model
	Email            string    `json:"email" gorm:"unique;not null"`
	Password         string    `json:"password" gorm:"-"`
	PasswordHash     []byte    `json:"-" gorm:"not null;"`
	ProfilePicture   []byte    `json:"profilePicture" gorm:"not null"`
	Name             string    `json:"name" gorm:"not null"`
	PhoneNumber      string    `json:"phoneNumber" gorm:"unique;not null"`
	Job              string    `json:"job" gorm:"not null"`
	Devices          []Device  `json:"devices" gorm:"constraint:OnDelete:CASCADE;"`
	MessagesSent     []Message `json:"messagesSent" gorm:"constraint:OnDelete:CASCADE;foreignKey:SenderUUID"`
	MessagesReceived []Message `json:"messagesReceived" gorm:"constraint:OnDelete:CASCADE;foreignKey:RecipientUUID"`
}

func RandomUser() *User {
	person := gofakeit.NewCrypto().Person()
	return &User{
		Password:       random.String()[:72],
		Email:          person.Contact.Email,
		ProfilePicture: []byte("JPEG IMAGE"),
		Name:           person.FirstName,
		PhoneNumber:    person.Contact.Phone,
		Job:            person.Job.Title,
	}
}

func (u *User) Claims() jwt.MapClaims {
	return jwt.MapClaims{
		"uuid": u.UUID.String(),
		"exp":  time.Now().Add(DefaultExpirationDelta).Unix(),
	}
}

func (u *User) FromClaims(m jwt.MapClaims) error {
	userUUID, ok := m["uuid"]
	if !ok {
		return fmt.Errorf("incomplete UUID")
	}
	u.UUID = uuid.FromStringOrNil(userUUID.(string))
	return nil
}

func (u *User) Authenticate(password string) bool {
	return u.PasswordHash != nil && bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) == nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	bErr := u.Model.BeforeSave(tx)
	if bErr != nil {
		return bErr
	}
	if len(u.Password) == 0 {
		return nil
	}
	var err error
	u.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(u.Password), DefaultPasswordCost)
	return err
}

var phoneRegex = regexp.MustCompile(`(?m)^\d{2,18}$`)

var (
	True  = true
	False = false
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	bErr := u.Model.BeforeSave(tx)
	if bErr != nil {
		return bErr
	}
	if len(u.PasswordHash) == 0 {
		return fmt.Errorf("password is empty")
	}
	if !phoneRegex.MatchString(u.PhoneNumber) {
		return fmt.Errorf("invalid phone")
	}
	_, pErr := mail.ParseAddress(u.Email)
	if pErr != nil {
		return pErr
	}
	return nil
}
