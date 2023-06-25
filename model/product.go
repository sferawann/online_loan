package model

import "time"

type Product struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Installment int64     `gorm:"not null" json:"installment"`
	Interest    float64   `gorm:"not null" json:"interest"`
	CreatedAt   time.Time `gorm:"->;not null; default:CURRENT_TIMESTAMP" json:"created_at"`
}
