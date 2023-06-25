package usecase

import (
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
)

type AcceptStatusUsecaseImpl struct {
	AcceptStatusRepo repository.AcceptStatusRepo
	TraRepo          repository.TraRepo
	ProRepo          repository.ProductRepo
}

// Delete implements AcceptStatusUsecase
func (u *AcceptStatusUsecaseImpl) Delete(id int64) (model.AcceptStatus, error) {
	return u.AcceptStatusRepo.Delete(id)
}

// Update implements AcceptStatusUsecase
func (u *AcceptStatusUsecaseImpl) Update(updatedAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	previousAcceptStatus, err := u.AcceptStatusRepo.FindById(updatedAcceptStatus.ID)
	if err != nil {
		return model.AcceptStatus{}, err
	}

	previousTraID := previousAcceptStatus.TransactionID
	previousStatus := previousAcceptStatus.Status
	previousCreatedAt := previousAcceptStatus.CreatedAt

	if updatedAcceptStatus.TransactionID == 0 {
		updatedAcceptStatus.TransactionID = previousTraID
	}
	if !updatedAcceptStatus.Status {
		updatedAcceptStatus.Status = previousStatus
	}
	if updatedAcceptStatus.CreatedAt.IsZero() {
		updatedAcceptStatus.CreatedAt = previousCreatedAt
	}

	return u.AcceptStatusRepo.Update(updatedAcceptStatus)
}

// FindAll implements AcceptStatusUsecase
func (u *AcceptStatusUsecaseImpl) FindAll() ([]model.AcceptStatus, error) {
	accstats, err := u.AcceptStatusRepo.FindAll()
	if err != nil {
		return nil, err
	}
	for i := range accstats {
		tra, err := u.TraRepo.FindById(accstats[i].TransactionID)
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
		accstats[i].Transaction = tra
	}
	return accstats, nil
}

// FindById implements AcceptStatusUsecase
func (u *AcceptStatusUsecaseImpl) FindById(id int64) (model.AcceptStatus, error) {
	accstat, err := u.AcceptStatusRepo.FindById(id)
	if err != nil {
		return model.AcceptStatus{}, err
	}
	tra, err := u.TraRepo.FindById(accstat.TransactionID)
	if err != nil {
		return model.AcceptStatus{}, err
	}

	pro, err := u.ProRepo.FindById(tra.ProductID)
	if err != nil {
		return model.AcceptStatus{}, err
	}
	tra.TotalInterest = (pro.Interest * tra.Amount) / 100
	tra.Total = tra.Amount + tra.TotalInterest
	tra.TotalMonth = tra.Total / float64(pro.Installment)
	accstat.Transaction = tra

	return accstat, nil
}

// Save implements AcceptStatusUsecase
func (u *AcceptStatusUsecaseImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	return u.AcceptStatusRepo.Save(newAcceptStatus)
}

func NewAcceptStatusUsecaseImpl(AcceptStatusRepo repository.AcceptStatusRepo, TraRepo repository.TraRepo, ProRepo repository.ProductRepo) AcceptStatusUsecase {
	return &AcceptStatusUsecaseImpl{
		AcceptStatusRepo: AcceptStatusRepo,
		TraRepo:          TraRepo,
		ProRepo:          ProRepo,
	}
}
