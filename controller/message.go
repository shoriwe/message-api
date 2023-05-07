package controller

import (
	"context"

	"firebase.google.com/go/messaging"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
)

func (c *Controller) SendNotificationWithFirebase(recipients []string, msg *models.Message) error {
	bResponse, sErr := c.Client.SendMulticast(
		context.Background(),
		&messaging.MulticastMessage{
			Tokens: recipients,
			Notification: &messaging.Notification{
				Title: msg.Title,
				Body:  msg.Body,
			},
		})
	if sErr != nil {
		return sErr
	}
	results := make([]models.MessageResponse, 0, len(bResponse.Responses))
	for _, response := range bResponse.Responses {
		var errS string
		if response.Error != nil {
			errS = response.Error.Error()
		}
		results = append(results, models.MessageResponse{
			MessageUUID: msg.UUID,
			Success:     response.Success,
			FirebaseId:  response.MessageID,
			Error:       errS,
		})
	}
	return c.DB.Create(results).Error
}

func (c *Controller) SendMessage(s *session.Session, msg *models.Message) error {
	m := &models.Message{
		SenderUUID:    s.UserUUID,
		RecipientUUID: msg.RecipientUUID,
		Title:         msg.Title,
		Body:          msg.Body,
	}
	cErr := c.DB.Create(m).Error
	if cErr != nil {
		return cErr
	}
	var recipients []string
	fErr := c.DB.
		Model(&models.Device{}).
		Select("firebase_token").
		Where("user_uuid = ?", msg.RecipientUUID).
		Find(&recipients).
		Error
	if fErr != nil {
		return fErr
	}
	return c.SendNotificationWithFirebase(recipients, m)
}
