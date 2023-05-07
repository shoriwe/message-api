package controller

import (
	"github.com/shoriwe/message-api/models"
)

func (c *Controller) Register(u *models.User) error {
	user := &models.User{
		Email:          u.Email,
		Password:       u.Password,
		ProfilePicture: u.ProfilePicture,
		Name:           u.Name,
		PhoneNumber:    u.PhoneNumber,
		Job:            u.Job,
	}
	return c.DB.Create(user).Error
}
