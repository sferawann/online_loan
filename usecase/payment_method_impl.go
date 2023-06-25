package usecase

import (
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
)

type PaymentMethodUsecaseImpl struct {
	PaymentMethodRepo repository.PaymentMethodRepo
}

// Delete implements PaymentMethodUsecase
func (u *PaymentMethodUsecaseImpl) Delete(id int64) (model.PaymentMethod, error) {
	return u.PaymentMethodRepo.Delete(id)
}

// FindAll implements PaymentMethodUsecase
func (u *PaymentMethodUsecaseImpl) FindAll() ([]model.PaymentMethod, error) {
	return u.PaymentMethodRepo.FindAll()
}

// FindById implements PaymentMethodUsecase
func (u *PaymentMethodUsecaseImpl) FindById(id int64) (model.PaymentMethod, error) {
	return u.PaymentMethodRepo.FindById(id)
}

// FindByPaymentMethodname implements PaymentMethodUsecase
func (u *PaymentMethodUsecaseImpl) FindByName(name string) (model.PaymentMethod, error) {
	return u.PaymentMethodRepo.FindByName(name)
}

// Save implements PaymentMethodUsecase
func (u *PaymentMethodUsecaseImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	return u.PaymentMethodRepo.Save(newPaymentMethod)
}

// Update implements PaymentMethodUsecase
func (u *PaymentMethodUsecaseImpl) Update(updatedPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	// Mendapatkan entitas PaymentMethod sebelumnya dari PaymentMethodRepo berdasarkan ID
	previousPaymentMethod, err := u.PaymentMethodRepo.FindById(updatedPaymentMethod.ID)
	if err != nil {
		return model.PaymentMethod{}, err
	}

	// Mengambil nilai created_at dari entitas sebelumnya
	createdAt := previousPaymentMethod.CreatedAt

	// Membuat updatedPaymentMethod dengan nilai created_at yang ada sebelumnya
	updatedPaymentMethod.CreatedAt = createdAt

	return u.PaymentMethodRepo.Update(updatedPaymentMethod)
}

func NewPaymentMethodUsecaseImpl(PaymentMethodRepo repository.PaymentMethodRepo) PaymentMethodUsecase {
	return &PaymentMethodUsecaseImpl{
		PaymentMethodRepo: PaymentMethodRepo,
	}
}
