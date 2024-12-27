package biz

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:50;not null;unique" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	Permissions []*Permission  `gorm:"many2many:role_permissions;" json:"permissions"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (*Role) TableName() string {
	return "t_roles"
}
