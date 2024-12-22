package biz

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`               // 用户id
	Username  string     `gorm:"type:varchar(50);not null;unique" json:"username"` // 用户名 唯一
	Password  string     `gorm:"size:255;not null" json:"-"`                       // 密码
	Email     string     `gorm:"size:100" json:"email"`                            // 邮箱
	Phone     string     `gorm:"size:11" json:"phone"`                             // 手机号
	Status    bool       `gorm:"default:true" json:"status"`                       // 状态
	Roles     []*Role    `gorm:"many2many:user_roles;" json:"roles"`               // 角色
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`                 // 创建时间
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`                 // 更新时间
	DeletedAt *time.Time `json:"-" gorm:"index;autoDeleteTime" `                   // 删除时间
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
