package model

import "time"

type User struct {
	ID          int64     `gorm:"primaryKey" form:"id" json:"id"`
	Username    string    `gorm:"unique;not null" form:"username" json:"username"`
	Password    string    `gorm:"not null" form:"password" json:"password"`
	KTP         string    `gorm:"not null" form:"ktp" json:"ktp"`
	Name        string    `gorm:"not null" form:"name" json:"name"`
	Address     string    `gorm:"not null" form:"address" json:"address"`
	PhoneNumber string    `gorm:"not null" form:"phone_number" json:"phone_number"`
	Limit       float64   `gorm:"not null" json:"limit"`
	RoleID      int64     `gorm:"column:id_role;not null"  json:"id_role"`
	CreatedAt   time.Time `gorm:"->;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	Role        Role      `json:"role"`
}
