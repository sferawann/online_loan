package repository

import "github.com/sferawann/pinjol/model"

type TraRepo interface {
	Save(newTra model.Transaction) (model.Transaction, error)
	FindById(id int64) (model.Transaction, error)
	FindAll() ([]model.Transaction, error)
	Update(updatedTra model.Transaction) (model.Transaction, error)
	Delete(id int64) (model.Transaction, error)
}
