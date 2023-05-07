package models

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/shoriwe/message-api/common/random"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const DefaultPasswordCost = 12

type User struct {
	Model
	Email            *string   `json:"email,omitempty" gorm:"unique;not null"`
	Password         *string   `json:"password,omitempty" gorm:"-"`
	PasswordHash     []byte    `json:"-" gorm:"not null;"`
	ProfilePicture   []byte    `json:"profilePicture,omitempty" gorm:"not null"`
	Name             *string   `json:"name,omitempty" gorm:"not null"`
	PhoneNumber      *string   `json:"phoneNumber,omitempty" gorm:"unique;not null"`
	Job              *string   `json:"job,omitempty" gorm:"not null"`
	Devices          []Device  `json:"devices,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	MessagesSent     []Message `json:"messagesSent,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:SenderUUID"`
	MessagesReceived []Message `json:"messagesReceived,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:RecipientUUID"`
}

func RandomUser() *User {
	password := random.String()[:72]
	person := gofakeit.NewCrypto().Person()
	return &User{
		Password:       &password,
		Email:          &person.Contact.Email,
		ProfilePicture: []byte("JPEG IMAGE"),
		Name:           &person.FirstName,
		PhoneNumber:    &person.Contact.Phone,
		Job:            &person.Job.Title,
	}
}

func (u *User) Authenticate(password string) bool {
	return u.PasswordHash != nil && bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) == nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	bErr := u.Model.BeforeSave(tx)
	if bErr != nil {
		return bErr
	}
	if u.Password != nil && len(*u.Password) > 0 {
		var err error
		u.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(*u.Password), DefaultPasswordCost)
		if err != nil {
			return err
		}
	}
	return nil
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
	if u.PhoneNumber == nil {
		return fmt.Errorf("phone is empty")
	}
	if !phoneRegex.MatchString(*u.PhoneNumber) {
		return fmt.Errorf("invalid phone")
	}
	if u.Email == nil {
		return fmt.Errorf("email field is empty")
	}
	_, pErr := mail.ParseAddress(*u.Email)
	if pErr != nil {
		return pErr
	}
	return nil
}
