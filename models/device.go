package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/message-api/common/random"
)

type Device struct {
	Model
	User          User      `json:"user" gorm:"foreignKey:UserUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserUUID      uuid.UUID `json:"userUUID" gorm:"uniqueIndex:idx_unique_token;not null;"`
	FirebaseToken string    `json:"firebaseToken" gorm:"uniqueIndex:idx_unique_token;not null;"`
}

func RandomDevice(u *User) *Device {
	return &Device{
		UserUUID:      u.UUID,
		FirebaseToken: random.String(),
	}
}
