package repository

import "github.com/sferawann/pinjol/model"

type ProductRepo interface {
	Save(newProduct model.Product) (model.Product, error)
	Update(updatedProduct model.Product) (model.Product, error)
	Delete(id int64) (model.Product, error)
	FindById(id int64) (model.Product, error)
	FindAll() ([]model.Product, error)
	FindByName(name string) (model.Product, error)
}
