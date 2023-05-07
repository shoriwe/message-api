package controller

import (
	"errors"

	"github.com/shoriwe/message-api/models"
	"gorm.io/gorm"
)

func (c *Controller) DownloadPicture(uuid string) ([]byte, error) {
	var user models.User
	err := c.DB.
		Where("uuid = ?", uuid).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUnauthorized
		}
		return nil, err
	}
	return user.ProfilePictureBytes, nil
}
