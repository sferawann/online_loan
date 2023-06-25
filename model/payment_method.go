package model

import "time"

type PaymentMethod struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `gorm:"->;not null; default:CURRENT_TIMESTAMP" json:"created_at"`
}
