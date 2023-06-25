package repository

import "github.com/sferawann/pinjol/model"

type PaymentRepo interface {
	Save(newPayment model.Payment) (model.Payment, error)
	Update(updatedPayment model.Payment) (model.Payment, error)
	Delete(id int64) (model.Payment, error)
	FindById(id int64) (model.Payment, error)
	FindAll() ([]model.Payment, error)
}
