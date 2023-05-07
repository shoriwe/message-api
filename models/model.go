package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Model struct {
	UUID      uuid.UUID  `json:"uuid" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (m *Model) BeforeSave(tx *gorm.DB) error {
	if m.UUID == uuid.Nil {
		m.UUID = uuid.NewV4()
	}
	return nil
}