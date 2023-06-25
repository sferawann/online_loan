package repository

import (
	"github.com/sferawann/pinjol/model"
	"gorm.io/gorm"
)

type TraRepoImpl struct {
	DB *gorm.DB
}

func NewTraRepoImpl(DB *gorm.DB) TraRepo {
	return &TraRepoImpl{DB: DB}
}

func (r *TraRepoImpl) Save(newTra model.Transaction) (model.Transaction, error) {
	result := r.DB.Create(&newTra)
	if result.Error != nil {
		return model.Transaction{}, result.Error
	}
	return newTra, nil
}

func (r *TraRepoImpl) FindById(id int64) (model.Transaction, error) {
	var Tra model.Transaction
	result := r.DB.Preload("User.Role").Preload("Product").First(&Tra, id)
	if result.Error != nil {
		return model.Transaction{}, result.Error
	}
	return Tra, nil
}

func (r *TraRepoImpl) FindAll() ([]model.Transaction, error) {
	var Tras []model.Transaction
	result := r.DB.Preload("User.Role").Preload("Product").Find(&Tras)
	if result.Error != nil {
		return nil, result.Error
	}
	return Tras, nil
}

func (r *TraRepoImpl) Update(updatedTra model.Transaction) (model.Transaction, error) {
	updateFields := make(map[string]interface{})

	// Tambahkan field dan nilai yang ingin diperbarui ke dalam map
	if updatedTra.UserID != 0 {
		updateFields["id_user"] = updatedTra.UserID
	}
	if updatedTra.ProductID != 0 {
		updateFields["id_product"] = updatedTra.ProductID
	}
	if updatedTra.Amount != 0 {
		updateFields["amount"] = updatedTra.Amount
	}
	if !updatedTra.Status {
		updateFields["status"] = updatedTra.Status
	}
	if updatedTra.CreatedAt.IsZero() {
		updateFields["created_at"] = updatedTra.CreatedAt
	}
	if updatedTra.DueDate.IsZero() {
		updateFields["due_date"] = updatedTra.DueDate
	}

	result := r.DB.Preload("User.Role").Model(&model.Transaction{}).Where("id = ?", updatedTra.ID).Updates(updateFields)
	if result.Error != nil {
		return model.Transaction{}, result.Error
	}
	return updatedTra, nil
}

func (r *TraRepoImpl) Delete(id int64) (model.Transaction, error) {
	var tra model.Transaction
	result := r.DB.First(&tra, id)
	if result.Error != nil {
		return model.Transaction{}, result.Error
	}
	err := r.DB.Delete(&tra).Error
	if err != nil {
		return model.Transaction{}, err
	}
	return tra, nil
}
