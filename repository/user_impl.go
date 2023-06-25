package repository

import (
	"github.com/sferawann/pinjol/model"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepoImpl(DB *gorm.DB) UserRepo {
	return &UserRepoImpl{DB: DB}
}

func (r *UserRepoImpl) Save(newUser model.User) (model.User, error) {
	result := r.DB.Create(&newUser)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return newUser, nil
}

func (r *UserRepoImpl) Update(updatedUser model.User) (model.User, error) {
	updateFields := make(map[string]interface{})

	// Tambahkan field dan nilai yang ingin diperbarui ke dalam map
	if updatedUser.Username != "" {
		updateFields["username"] = updatedUser.Username
	}
	if updatedUser.Password != "" {
		updateFields["password"] = updatedUser.Password
	}
	if updatedUser.KTP != "" {
		updateFields["ktp"] = updatedUser.KTP
	}
	if updatedUser.Name != "" {
		updateFields["name"] = updatedUser.Name
	}
	if updatedUser.Address != "" {
		updateFields["address"] = updatedUser.Address
	}
	if updatedUser.PhoneNumber != "" {
		updateFields["phone_number"] = updatedUser.PhoneNumber
	}
	if updatedUser.Limit != 0 {
		updateFields["limit"] = updatedUser.Limit
	}
	if updatedUser.RoleID != 0 {
		updateFields["id_role"] = updatedUser.RoleID
	}
	if updatedUser.CreatedAt.IsZero() {
		updateFields["created_at"] = updatedUser.CreatedAt
	}

	result := r.DB.Preload("Role").Model(&model.User{}).Where("id = ?", updatedUser.ID).Updates(updateFields)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return updatedUser, nil
}

func (r *UserRepoImpl) Delete(id int64) (model.User, error) {
	var user model.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	err := r.DB.Delete(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepoImpl) FindById(id int64) (model.User, error) {
	var user model.User
	result := r.DB.Preload("Role").First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepoImpl) FindAll() ([]model.User, error) {
	var users []model.User
	result := r.DB.Preload("Role").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepoImpl) FindByUsername(username string) (model.User, error) {
	var user model.User
	result := r.DB.Preload("Role").Where("username = ?", username).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
