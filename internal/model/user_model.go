package model

type Users struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (*Users) TableName() string {
	return "t_users"
}
