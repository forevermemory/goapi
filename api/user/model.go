package user

import (
	"time"
)

type User struct {
	ID	int	`json:"id" gorm:"column:id;primary_key;auto_increment;"`

	Username	string	`json:"username" gorm:"column:username"`
	Password	string	`json:"password" gorm:"column:password"`
	Email	string	`json:"email" gorm:"column:email"`
	Blocked	int	`json:"blocked" gorm:"column:blocked"`
	LoginAt	time.Time	`json:"login_at" gorm:"column:login_at"`

	CreatedAt	time.Time	`json:"created_at" gorm:"column:created_at"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"column:updated_at"`
	DeletedAt	*time.Time	`json:"-" gorm:"column:deleted_at"`
}


