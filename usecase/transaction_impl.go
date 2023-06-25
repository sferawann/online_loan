package usecase

import (
	"time"

	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
)

type TraUsecaseImpl struct {
	TraRepo  repository.TraRepo
	UserRepo repository.UserRepo
	ProRepo  repository.ProductRepo
}

// FindAll implements TraUsecase
func (u *TraUsecaseImpl) FindAll() ([]model.Transaction, error) {
	tras, err := u.TraRepo.FindAll()
	if err != nil {
		return nil, err
	}

	for i := range tras {
		product := model.Product{}
		product, err := u.ProRepo.FindById(tras[i].ProductID)
		if err != nil {
			return nil, err
		}
		tras[i].TotalInterest = (product.Interest * tras[i].Amount) / 100
		tras[i].Total = tras[i].Amount + tras[i].TotalInterest
		tras[i].TotalMonth = tras[i].Total / float64(product.Installment)
	}

	return tras, nil

}

// FindById implements TraUsecase
func (u *TraUsecaseImpl) FindById(id int64) (model.Transaction, error) {
	tra, err := u.TraRepo.FindById(id)
	if err != nil {
		return model.Transaction{}, err
	}
	product := model.Product{}
	product, err = u.ProRepo.FindById(tra.ProductID)
	if err != nil {
		return model.Transaction{}, err
	}
	tra.TotalInterest = (product.Interest * tra.Amount) / 100
	tra.Total = tra.Amount + tra.TotalInterest
	tra.TotalMonth = tra.Total / float64(product.Installment)

	return tra, nil
}

// Save implements TraUsecase
func (u *TraUsecaseImpl) Save(newTra model.Transaction) (model.Transaction, error) {
	newTra.DueDate = time.Now().AddDate(0, 1, 0)

	user, err := u.UserRepo.FindById(newTra.UserID)
	if err != nil {
		return model.Transaction{}, err
	}
	newTra.User = user

	product, err := u.ProRepo.FindById(newTra.ProductID)
	if err != nil {
		return model.Transaction{}, err
	}
	newTra.Product = product

	newTra.TotalInterest = (product.Interest * newTra.Amount) / 100
	newTra.Total = newTra.Amount + newTra.TotalInterest
	newTra.TotalMonth = newTra.Total / float64(product.Installment)

	return u.TraRepo.Save(newTra)
}

func (u *TraUsecaseImpl) Update(UpdatedTra model.Transaction) (model.Transaction, error) {
	previousTra, err := u.TraRepo.FindById(UpdatedTra.ID)
	if err != nil {
		return model.Transaction{}, err
	}

	previousUserID := previousTra.UserID
	previousProductID := previousTra.ProductID
	previousAmount := previousTra.Amount
	previousStatus := previousTra.Status
	previousCreatedAt := previousTra.CreatedAt
	previousDueDate := previousTra.DueDate

	if UpdatedTra.UserID == 0 {
		UpdatedTra.UserID = previousUserID
	}
	if UpdatedTra.ProductID == 0 {
		UpdatedTra.ProductID = previousProductID
	}
	if UpdatedTra.Amount == 0 {
		UpdatedTra.Amount = previousAmount
	}
	if !UpdatedTra.Status {
		UpdatedTra.Status = previousStatus
	}
	if UpdatedTra.CreatedAt.IsZero() {
		UpdatedTra.CreatedAt = previousCreatedAt
	}
	if UpdatedTra.DueDate.IsZero() {
		UpdatedTra.DueDate = previousDueDate
	}

	return u.TraRepo.Update(UpdatedTra)
}

func (u *TraUsecaseImpl) Delete(id int64) (model.Transaction, error) {
	return u.TraRepo.Delete(id)
}

func NewTraUsecaseImpl(TraRepo repository.TraRepo, UserRepo repository.UserRepo, ProRepo repository.ProductRepo) TraUsecase {
	return &TraUsecaseImpl{
		TraRepo:  TraRepo,
		UserRepo: UserRepo,
		ProRepo:  ProRepo,
	}
}
