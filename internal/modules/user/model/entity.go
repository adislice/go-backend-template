package user_model

import "time"

// User model
type User struct {
	ID        string    `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
