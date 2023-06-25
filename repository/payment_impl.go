package repository

import (
	"github.com/sferawann/pinjol/model"
	"gorm.io/gorm"
)

type PaymentRepoImpl struct {
	DB *gorm.DB
}

func NewPaymentRepoImpl(DB *gorm.DB) PaymentRepo {
	return &PaymentRepoImpl{DB: DB}
}

func (r *PaymentRepoImpl) Save(newPayment model.Payment) (model.Payment, error) {
	result := r.DB.Create(&newPayment)
	if result.Error != nil {
		return model.Payment{}, result.Error
	}
	return newPayment, nil
}

func (r *PaymentRepoImpl) Update(updatedPayment model.Payment) (model.Payment, error) {
	updateFields := make(map[string]interface{})

	// Tambahkan field dan nilai yang ingin diperbarui ke dalam map
	if updatedPayment.TransactionID != 0 {
		updateFields["id_transaction"] = updatedPayment.TransactionID
	}
	if updatedPayment.PaymentAmount != 0 {
		updateFields["payment_amount"] = updatedPayment.PaymentAmount
	}
	if updatedPayment.PaymentMethodID != 0 {
		updateFields["id_payment_method"] = updatedPayment.PaymentMethodID
	}
	if updatedPayment.PaymentDate.IsZero() {
		updateFields["payment_date"] = updatedPayment.PaymentDate
	}

	result := r.DB.Preload("Transaction").Preload("PaymentMethod").Model(&model.Payment{}).Where("id = ?", updatedPayment.ID).Updates(updateFields)
	if result.Error != nil {
		return model.Payment{}, result.Error
	}
	return updatedPayment, nil
}

func (r *PaymentRepoImpl) Delete(id int64) (model.Payment, error) {
	var Payment model.Payment
	result := r.DB.First(&Payment, id)
	if result.Error != nil {
		return model.Payment{}, result.Error
	}
	err := r.DB.Delete(&Payment).Error
	if err != nil {
		return model.Payment{}, err
	}
	return Payment, nil
}

func (r *PaymentRepoImpl) FindById(id int64) (model.Payment, error) {
	var Payment model.Payment
	result := r.DB.Preload("Transaction.User.Role").Preload("Transaction.Product").Preload("PaymentMethod").First(&Payment, id)
	if result.Error != nil {
		return model.Payment{}, result.Error
	}
	return Payment, nil
}

func (r *PaymentRepoImpl) FindAll() ([]model.Payment, error) {
	var Payments []model.Payment
	result := r.DB.Preload("Transaction.User.Role").Preload("Transaction.Product").Preload("PaymentMethod").Find(&Payments)
	if result.Error != nil {
		return nil, result.Error
	}
	return Payments, nil
}
