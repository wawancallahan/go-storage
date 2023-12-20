package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (m *Book) TableName() string {
	return "books"
}
