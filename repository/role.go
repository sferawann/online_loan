package repository

import "github.com/sferawann/pinjol/model"

type RoleRepo interface {
	Save(newRole model.Role) (model.Role, error)
	Update(updatedRole model.Role) (model.Role, error)
	Delete(id int64) (model.Role, error)
	FindById(id int64) (model.Role, error)
	FindAll() ([]model.Role, error)
	FindByName(name string) (model.Role, error)
}
