package model

import "time"

type Transaction struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `gorm:"column:id_user;not null"  json:"id_user"`
	User      User      `json:"user"`
	ProductID int64     `gorm:"column:product_id;not null" json:"id_product"`
	Product   Product   `json:"product"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Status    bool      `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"->;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	DueDate   time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"due_date"`

	TotalInterest float64 `gorm:"-" json:"total_interest"`
	TotalMonth    float64 `gorm:"-" json:"total_month"`
	Total         float64 `gorm:"-" json:"total"`
}
