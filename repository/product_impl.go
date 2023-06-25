package repository

import (
	"github.com/sferawann/pinjol/model"
	"gorm.io/gorm"
)

type ProductRepoImpl struct {
	DB *gorm.DB
}

func NewProductRepoImpl(DB *gorm.DB) ProductRepo {
	return &ProductRepoImpl{DB: DB}
}

func (r *ProductRepoImpl) Save(newProduct model.Product) (model.Product, error) {
	result := r.DB.Create(&newProduct)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return newProduct, nil
}

func (r *ProductRepoImpl) Update(updatedProduct model.Product) (model.Product, error) {
	updateFields := make(map[string]interface{})

	// Tambahkan field dan nilai yang ingin diperbarui ke dalam map
	if updatedProduct.Name != "" {
		updateFields["name"] = updatedProduct.Name
	}
	if updatedProduct.Installment != 0 {
		updateFields["installment"] = updatedProduct.Installment
	}
	if updatedProduct.Interest != 0 {
		updateFields["interest"] = updatedProduct.Interest
	}
	if updatedProduct.CreatedAt.IsZero() {
		updateFields["created_at"] = updatedProduct.CreatedAt
	}

	result := r.DB.Model(&model.Product{}).Where("id = ?", updatedProduct.ID).Updates(updateFields)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return updatedProduct, nil
}

func (r *ProductRepoImpl) Delete(id int64) (model.Product, error) {
	var Product model.Product
	result := r.DB.First(&Product, id)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	err := r.DB.Delete(&Product).Error
	if err != nil {
		return model.Product{}, err
	}
	return Product, nil
}

func (r *ProductRepoImpl) FindById(id int64) (model.Product, error) {
	var Product model.Product
	result := r.DB.First(&Product, id)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return Product, nil
}

func (r *ProductRepoImpl) FindAll() ([]model.Product, error) {
	var Products []model.Product
	result := r.DB.Find(&Products)
	if result.Error != nil {
		return nil, result.Error
	}
	return Products, nil
}

func (r *ProductRepoImpl) FindByName(name string) (model.Product, error) {
	var Product model.Product
	result := r.DB.Where("name = ?", name).First(&Product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return Product, nil
}
