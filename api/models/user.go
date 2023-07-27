package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"size:100;not null;unique" json:"email" validate:"required,email,max=100"`
	Password  string    `gorm:"not null" json:"-" validate:"required,min=6"`
	Name      *string   `gorm:"size:20;unique" json:"name" validate:"omitempty,max=20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "tb_user"
}
