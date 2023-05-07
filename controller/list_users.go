package controller

import (
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
)

func (c *Controller) ListUsers(s *session.Session, filter *Filter[models.User]) ([]models.User, error) {
	filter.Page--
	if filter.Page < 0 {
		filter.Page = 0
	}
	if filter.PageSize <= 0 || filter.PageSize > 100 {
		filter.PageSize = 100
	}
	var result []models.User
	query := c.DB.
		Where("uuid != ?", s.UserUUID)
	if filter.Target.Email != nil {
		query = query.Where("email LIKE ?", *filter.Target.Email)
	}
	if filter.Target.Name != nil {
		query = query.Where("name LIKE ?", *filter.Target.Name)
	}
	if filter.Target.Job != nil {
		query = query.Where("job LIKE ?", *filter.Target.Job)
	}
	if filter.Target.PhoneNumber != nil {
		query = query.Where("phone_number LIKE ?", *filter.Target.PhoneNumber)
	}
	fErr := query.
		Limit(filter.PageSize).
		Offset(filter.Page).
		Find(&result).
		Error
	return result, fErr
}
