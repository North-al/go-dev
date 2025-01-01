package biz

import "gorm.io/gorm"

type Menu struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	// 父级菜单、vue文件路径、路由路径、路由name、重定向路径 { meta: { title, icon, hidden, keepAlive,  } }
	ParentID uint64 `json:"parent_id"`
	// 页面文件路径
	PageFilePath string `json:"page_file_path"`
	// 路由路径
	RoutePath string `json:"route_path"`
	// 路由name
	RouteName string `json:"route_name"`
	// 重定向路径
	Redirect string `json:"redirect"`
	// 路由菜单标题
	Title string `json:"title"`
	// 路由菜单图标
	Icon string `json:"icon"`
	// 路由菜单是否隐藏
	Hidden bool `json:"hidden"`
	// 路由菜单是否缓存
	KeepAlive bool `json:"keep_alive"`
	// 路由菜单排序
	Sort int `json:"sort"`

	CreatedAt *LocalTime     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *LocalTime     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;autoDeleteTime"`
}

func (*Menu) TableName() string {
	return "t_menus"
}
