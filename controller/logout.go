package controller

import (
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
)

func (c *Controller) Logout(s *session.Session) error {
	query := c.DB.
		Unscoped().
		Where("user_uuid = ?", s.UserUUID).
		Where("firebase_token = ?", s.FirebaseToken).
		Delete(&models.Device{})
	if err := query.Error; err != nil {
		return err
	}
	if query.RowsAffected == 0 {
		return ErrUnauthorized
	}
	return nil
}
