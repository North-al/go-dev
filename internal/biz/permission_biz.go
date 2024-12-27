package biz

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:50;not null;unique" json:"name"`
	Code        string         `gorm:"size:100;not null;unique" json:"code"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (*Permission) TableName() string {
	return "t_permissions"
}
