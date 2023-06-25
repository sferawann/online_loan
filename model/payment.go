package model

import "time"

type Payment struct {
	ID              int64         `gorm:"primaryKey" json:"id"`
	TransactionID   int64         `gorm:"column:id_transaction" json:"id_transaction"`
	Transaction     Transaction   `json:"transaction"`
	PaymentAmount   float64       `gorm:"not null" json:"payment_amount"`
	PaymentMethodID int64         `gorm:"not null" json:"id_payment_method"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	PaymentDate     time.Time     `gorm:"->;not null; default:CURRENT_TIMESTAMP" json:"payment_date"`

	NextInstallment float64 `gorm:"-" json:"next_installment"`
}
