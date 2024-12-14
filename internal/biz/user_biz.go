package biz

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string     `gorm:"size:50;not null;unique" json:"username"`
	Password  string     `gorm:"size:255;not null" json:"-"`
	Email     string     `gorm:"size:100" json:"email"`
	Phone     string     `gorm:"size:15" json:"phone"`
	Status    bool       `gorm:"default:true" json:"status"`
	Roles     []*Role    `gorm:"many2many:user_roles;" json:"roles"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

func (*Users) TableName() string {
	return "t_users"
}

// AfterFind 钩子函数
func (u *Users) AfterFind(tx *gorm.DB) (err error) {
	if u.Roles == nil {
		u.Roles = []*Role{}
	}
	return nil
}
