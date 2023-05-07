package controller

import (
	"errors"

	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"gorm.io/gorm"
)

type Login struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirebaseToken string `json:"firebaseToken"`
}

func (c *Controller) Login(req *Login) (*models.User, string, error) {
	user := new(models.User)
	fErr := c.DB.
		Where("email = ?", req.Email).
		First(&user).
		Error
	if fErr != nil {
		if errors.Is(fErr, gorm.ErrRecordNotFound) {
			return nil, "", ErrUnauthorized
		}
		return nil, "", fErr
	}
	if !user.Authenticate(req.Password) {
		return nil, "", ErrUnauthorized
	}
	deviceErr := c.DB.Create(&models.Device{
		UserUUID:      user.UUID,
		FirebaseToken: req.FirebaseToken,
	}).Error
	return user, c.JWT.New(session.NewSession(user.UUID, req.FirebaseToken)), deviceErr
}
