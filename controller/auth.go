package controller

import (
	"errors"

	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
	"gorm.io/gorm"
)

func (c *Controller) Authenticate(s *session.Session) (*models.User, error) {
	var device models.Device
	fErr := c.DB.
		Preload("User").
		Where("user_uuid = ?", s.UserUUID).
		Where("firebase_token = ?", s.FirebaseToken).
		First(&device).
		Error
	if fErr != nil {
		if errors.Is(fErr, gorm.ErrRecordNotFound) {
			return nil, ErrUnauthorized
		}
		return nil, fErr
	}
	return &device.User, nil
}
