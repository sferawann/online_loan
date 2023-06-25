package usecase

import (
	"time"

	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
)

type ProductUsecaseImpl struct {
	ProductRepo repository.ProductRepo
}

// Delete implements ProductUsecase
func (u *ProductUsecaseImpl) Delete(id int64) (model.Product, error) {
	return u.ProductRepo.Delete(id)
}

// FindAll implements ProductUsecase
func (u *ProductUsecaseImpl) FindAll() ([]model.Product, error) {
	return u.ProductRepo.FindAll()
}

// FindById implements ProductUsecase
func (u *ProductUsecaseImpl) FindById(id int64) (model.Product, error) {
	return u.ProductRepo.FindById(id)
}

// FindByProductname implements ProductUsecase
func (u *ProductUsecaseImpl) FindByName(name string) (model.Product, error) {
	return u.ProductRepo.FindByName(name)
}

// Save implements ProductUsecase
func (u *ProductUsecaseImpl) Save(newProduct model.Product) (model.Product, error) {
	return u.ProductRepo.Save(newProduct)
}

// Update implements ProductUsecase
func (u *ProductUsecaseImpl) Update(updatedProduct model.Product) (model.Product, error) {
	// Mendapatkan entitas Product sebelumnya dari ProductRepo berdasarkan ID
	previousProduct, err := u.ProductRepo.FindById(updatedProduct.ID)
	if err != nil {
		return model.Product{}, err
	}

	// Mengambil nilai created_at dari entitas sebelumnya
	previousName := previousProduct.Name
	previousInstallment := previousProduct.Installment
	previousInterest := previousProduct.Interest
	previousCreatedAt := previousProduct.CreatedAt

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedProduct.Name == "" {
		updatedProduct.Name = previousName
	}
	if updatedProduct.Installment == 0 {
		updatedProduct.Installment = previousInstallment
	}
	if updatedProduct.Interest == 0 {
		updatedProduct.Interest = previousInterest
	}
	if updatedProduct.CreatedAt == (time.Time{}) {
		updatedProduct.CreatedAt = previousCreatedAt
	}

	return u.ProductRepo.Update(updatedProduct)
}

func NewProductUsecaseImpl(ProductRepo repository.ProductRepo) ProductUsecase {
	return &ProductUsecaseImpl{
		ProductRepo: ProductRepo,
	}
}
