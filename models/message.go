package models

import (
	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/common/random"
)

type Message struct {
	Model
	Sender        User      `json:"sender" gorm:"foreignKey:SenderUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SenderUUID    uuid.UUID `json:"senderUUID" gorm:"not null;"`
	Recipient     User      `json:"recipient" gorm:"foreignKey:RecipientUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RecipientUUID uuid.UUID `json:"recipientUUID" gorm:"not null;"`
	Title         string    `json:"title" gorm:"not null"`
	Body          string    `json:"body" gorm:"not null"`
}

func RandomMessage(sender, recipient *User) *Message {
	return &Message{
		SenderUUID:    sender.UUID,
		RecipientUUID: recipient.UUID,
		Title:         gofakeit.NewCrypto().Name(),
		Body:          random.String(),
	}
}

type MessageResponse struct {
	Model
	Message     Message   `json:"message" gorm:"foreignKey:MessageUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MessageUUID uuid.UUID `json:"messageUUID" gorm:"not null;"`
	Success     bool      `json:"success" gorm:"not null"`
	FirebaseId  string    `json:"firebaseId" gorm:"not null"`
	Error       string    `json:"error" gorm:"not null"`
}
