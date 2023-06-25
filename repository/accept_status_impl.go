package repository

import (
	"github.com/sferawann/pinjol/model"
	"gorm.io/gorm"
)

type AcceptStatusRepoImpl struct {
	DB *gorm.DB
}

func NewAcceptStatusRepoImpl(DB *gorm.DB) AcceptStatusRepo {
	return &AcceptStatusRepoImpl{DB: DB}
}

func (r *AcceptStatusRepoImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	result := r.DB.Create(&newAcceptStatus)
	if result.Error != nil {
		return model.AcceptStatus{}, result.Error
	}
	return newAcceptStatus, nil
}

func (r *AcceptStatusRepoImpl) FindById(id int64) (model.AcceptStatus, error) {
	var AcceptStatus model.AcceptStatus
	result := r.DB.Preload("Transaction").First(&AcceptStatus, id)
	if result.Error != nil {
		return model.AcceptStatus{}, result.Error
	}
	return AcceptStatus, nil
}

func (r *AcceptStatusRepoImpl) FindAll() ([]model.AcceptStatus, error) {
	var AcceptStatuss []model.AcceptStatus
	result := r.DB.Preload("Transaction.User.Role").Preload("Transaction.Product").Find(&AcceptStatuss)
	if result.Error != nil {
		return nil, result.Error
	}
	return AcceptStatuss, nil
}

func (r *AcceptStatusRepoImpl) Update(updatedAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	result := r.DB.Save(&updatedAcceptStatus)
	if result.Error != nil {
		return model.AcceptStatus{}, result.Error
	}
	return updatedAcceptStatus, nil
}

func (r *AcceptStatusRepoImpl) Delete(id int64) (model.AcceptStatus, error) {
	var AcceptStatus model.AcceptStatus
	result := r.DB.First(&AcceptStatus, id)
	if result.Error != nil {
		return model.AcceptStatus{}, result.Error
	}
	err := r.DB.Delete(&AcceptStatus).Error
	if err != nil {
		return model.AcceptStatus{}, err
	}
	return AcceptStatus, nil
}