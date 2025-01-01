package biz

import (
	"gorm.io/gorm"
)

type Role struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:50;not null;unique" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	Menus       []*Menu        `gorm:"many2many:role_menus;" json:"menus"`
	CreatedAt   *LocalTime     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *LocalTime     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index;autoDeleteTime" json:"-"`
}

func (*Role) TableName() string {
	return "t_roles"
}
