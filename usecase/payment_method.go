package usecase

import "github.com/sferawann/pinjol/model"

type PaymentMethodUsecase interface {
	Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Update(updatedPaymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	Delete(id int64) (model.PaymentMethod, error)
	FindById(id int64) (model.PaymentMethod, error)
	FindAll() ([]model.PaymentMethod, error)
	FindByName(name string) (model.PaymentMethod, error)
}
