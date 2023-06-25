package usecase

import "github.com/sferawann/pinjol/model"

type TraUsecase interface {
	Save(newTra model.Transaction) (model.Transaction, error)
	FindById(id int64) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
	Update(UpdatedTra model.Transaction) (model.Transaction, error)
	Delete(id int64) (model.Transaction, error)
}
