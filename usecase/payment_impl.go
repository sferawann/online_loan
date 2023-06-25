package usecase

import (
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
)

type PaymentUsecaseImpl struct {
	PaymentRepo   repository.PaymentRepo
	TraRepo       repository.TraRepo
	ProRepo       repository.ProductRepo
	PayMethodRepo repository.PaymentMethodRepo
	UserRepo      repository.UserRepo
}

// Delete implements PaymentUsecase
func (u *PaymentUsecaseImpl) Delete(id int64) (model.Payment, error) {
	return u.PaymentRepo.Delete(id)
}

// FindAll implements PaymentUsecase
func (u *PaymentUsecaseImpl) FindAll() ([]model.Payment, error) {
	payments, err := u.PaymentRepo.FindAll()
	if err != nil {
		return nil, err
	}

	for i := range payments {
		tra, err := u.TraRepo.FindById(payments[i].TransactionID)
		if err != nil {
			return nil, err
		}

		pro, err := u.ProRepo.FindById(tra.ProductID)
		if err != nil {
			return nil, err
		}
		tra.TotalInterest = (pro.Interest * tra.Amount) / 100
		tra.Total = tra.Amount + tra.TotalInterest
		tra.TotalMonth = tra.Total / float64(pro.Installment)
		payments[i].NextInstallment = tra.Total - tra.TotalMonth
		payments[i].Transaction = tra
	}
	return payments, nil
}

// FindById implements PaymentUsecase
func (u *PaymentUsecaseImpl) FindById(id int64) (model.Payment, error) {
	payment, err := u.PaymentRepo.FindById(id)
	if err != nil {
		return model.Payment{}, err
	}
	tra, err := u.TraRepo.FindById(payment.TransactionID)
	if err != nil {
		return model.Payment{}, err
	}
	pro, err := u.ProRepo.FindById(tra.ProductID)
	if err != nil {
		return model.Payment{}, err
	}

	tra.TotalInterest = (pro.Interest * tra.Amount) / 100
	tra.Total = tra.Amount + tra.TotalInterest
	tra.TotalMonth = tra.Total / float64(pro.Installment)
	payment.NextInstallment = tra.Total - tra.TotalMonth
	payment.Transaction = tra

	return payment, nil
}

// Save implements PaymentUsecase
func (u *PaymentUsecaseImpl) Save(newPayment model.Payment) (model.Payment, error) {
	tra, err := u.TraRepo.FindById(newPayment.TransactionID)
	if err != nil {
		return model.Payment{}, err
	}
	newPayment.Transaction = tra

	user, err := u.UserRepo.FindById(tra.UserID)
	if err != nil {
		return model.Payment{}, err
	}
	tra.User = user

	product, err := u.ProRepo.FindById(tra.ProductID)
	if err != nil {
		return model.Payment{}, err
	}
	tra.Product = product

	newPayment.Transaction.TotalInterest = (product.Interest * tra.Amount) / 100
	newPayment.Transaction.Total = newPayment.Transaction.Amount + newPayment.Transaction.TotalInterest
	newPayment.Transaction.TotalMonth = newPayment.Transaction.Total / float64(product.Installment)
	newPayment.NextInstallment = newPayment.Transaction.Total - newPayment.Transaction.TotalMonth

	payMethod, err := u.PayMethodRepo.FindById(newPayment.PaymentMethodID)
	if err != nil {
		return model.Payment{}, err
	}
	newPayment.PaymentMethod = payMethod

	return u.PaymentRepo.Save(newPayment)
}

// Update implements PaymentUsecase
func (u *PaymentUsecaseImpl) Update(updatedPayment model.Payment) (model.Payment, error) {

	// Mendapatkan entitas Payment sebelumnya dari PaymentRepo berdasarkan ID
	previousPayment, err := u.PaymentRepo.FindById(updatedPayment.ID)
	if err != nil {
		return model.Payment{}, err
	}

	// Mengambil nilai-nilai field dari entitas sebelumnya
	previousTraID := previousPayment.TransactionID
	previousPaymentAmount := previousPayment.PaymentAmount
	previousPaymentMethodID := previousPayment.PaymentMethodID
	previousPaymentDate := previousPayment.PaymentDate

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedPayment.TransactionID == 0 {
		updatedPayment.TransactionID = previousTraID
	}
	if updatedPayment.PaymentAmount == 0 {
		updatedPayment.PaymentAmount = previousPaymentAmount
	}
	if updatedPayment.PaymentMethodID == 0 {
		updatedPayment.PaymentMethodID = previousPaymentMethodID
	}
	if updatedPayment.PaymentDate.IsZero() {
		updatedPayment.PaymentDate = previousPaymentDate
	}
	return u.PaymentRepo.Update(updatedPayment)
}

func NewPaymentUsecaseImpl(PaymentRepo repository.PaymentRepo, TraRepo repository.TraRepo, ProRepo repository.ProductRepo, PayMethodRepo repository.PaymentMethodRepo, UserRepo repository.UserRepo) PaymentUsecase {
	return &PaymentUsecaseImpl{
		PaymentRepo:   PaymentRepo,
		TraRepo:       TraRepo,
		ProRepo:       ProRepo,
		PayMethodRepo: PayMethodRepo,
		UserRepo:      UserRepo,
	}
}
