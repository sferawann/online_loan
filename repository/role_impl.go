package repository

import (
	"github.com/sferawann/pinjol/model"
	"gorm.io/gorm"
)

type RoleRepoImpl struct {
	DB *gorm.DB
}

func NewRoleRepoImpl(DB *gorm.DB) RoleRepo {
	return &RoleRepoImpl{DB: DB}
}

func (r *RoleRepoImpl) Save(newRole model.Role) (model.Role, error) {
	result := r.DB.Create(&newRole)
	if result.Error != nil {
		return model.Role{}, result.Error
	}
	return newRole, nil
}

func (r *RoleRepoImpl) Update(updatedRole model.Role) (model.Role, error) {
	result := r.DB.Save(&updatedRole)
	if result.Error != nil {
		return model.Role{}, result.Error
	}
	return updatedRole, nil
}

func (r *RoleRepoImpl) Delete(id int64) (model.Role, error) {
	var Role model.Role
	result := r.DB.First(&Role, id)
	if result.Error != nil {
		return model.Role{}, result.Error
	}
	err := r.DB.Delete(&Role).Error
	if err != nil {
		return model.Role{}, err
	}
	return Role, nil
}

func (r *RoleRepoImpl) FindById(id int64) (model.Role, error) {
	var Role model.Role
	result := r.DB.First(&Role, id)
	if result.Error != nil {
		return model.Role{}, result.Error
	}
	return Role, nil
}

func (r *RoleRepoImpl) FindAll() ([]model.Role, error) {
	var Roles []model.Role
	result := r.DB.Find(&Roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return Roles, nil
}

func (r *RoleRepoImpl) FindByName(name string) (model.Role, error) {
	var Role model.Role
	result := r.DB.Where("name = ?", name).First(&Role)
	if result.Error != nil {
		return model.Role{}, result.Error
	}
	return Role, nil
}
