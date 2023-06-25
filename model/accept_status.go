package model

import "time"

type AcceptStatus struct {
	ID            int64       `gorm:"primaryKey" json:"id"`
	TransactionID int64       `gorm:"column:id_transaction;not null" json:"id_transaction"`
	Transaction   Transaction `json:"transaction"`
	Status        bool        `gorm:"not null" json:"status"`
	CreatedAt     time.Time   `gorm:"->;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
