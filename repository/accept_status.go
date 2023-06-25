package repository

import "github.com/sferawann/pinjol/model"

type AcceptStatusRepo interface {
	Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
	FindById(id int64) (model.AcceptStatus, error)
	FindAll() ([]model.AcceptStatus, error)
	Update(updatedAcceptStatus model.AcceptStatus) (model.AcceptStatus, error)
	Delete(id int64) (model.AcceptStatus, error)
}
