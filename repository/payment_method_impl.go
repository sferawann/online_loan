package repository

import (
	"github.com/sferawann/pinjol/model"
	"gorm.io/gorm"
)

type PaymentMethodRepoImpl struct {
	DB *gorm.DB
}

func NewPaymentMethodRepoImpl(DB *gorm.DB) PaymentMethodRepo {
	return &PaymentMethodRepoImpl{DB: DB}
}

func (r *PaymentMethodRepoImpl) Save(newPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	result := r.DB.Create(&newPaymentMethod)
	if result.Error != nil {
		return model.PaymentMethod{}, result.Error
	}
	return newPaymentMethod, nil
}

func (r *PaymentMethodRepoImpl) Update(updatedPaymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	result := r.DB.Save(&updatedPaymentMethod)
	if result.Error != nil {
		return model.PaymentMethod{}, result.Error
	}
	return updatedPaymentMethod, nil
}

func (r *PaymentMethodRepoImpl) Delete(id int64) (model.PaymentMethod, error) {
	var PaymentMethod model.PaymentMethod
	result := r.DB.First(&PaymentMethod, id)
	if result.Error != nil {
		return model.PaymentMethod{}, result.Error
	}
	err := r.DB.Delete(&PaymentMethod).Error
	if err != nil {
		return model.PaymentMethod{}, err
	}
	return PaymentMethod, nil
}

func (r *PaymentMethodRepoImpl) FindById(id int64) (model.PaymentMethod, error) {
	var PaymentMethod model.PaymentMethod
	result := r.DB.First(&PaymentMethod, id)
	if result.Error != nil {
		return model.PaymentMethod{}, result.Error
	}
	return PaymentMethod, nil
}

func (r *PaymentMethodRepoImpl) FindAll() ([]model.PaymentMethod, error) {
	var PaymentMethods []model.PaymentMethod
	result := r.DB.Find(&PaymentMethods)
	if result.Error != nil {
		return nil, result.Error
	}
	return PaymentMethods, nil
}

func (r *PaymentMethodRepoImpl) FindByName(name string) (model.PaymentMethod, error) {
	var PaymentMethod model.PaymentMethod
	result := r.DB.Where("name = ?", name).First(&PaymentMethod)
	if result.Error != nil {
		return model.PaymentMethod{}, result.Error
	}
	return PaymentMethod, nil
}
