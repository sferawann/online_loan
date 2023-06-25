package usecase

import (
	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
)

type RoleUsecaseImpl struct {
	RoleRepo repository.RoleRepo
}

// Delete implements RoleUsecase
func (u *RoleUsecaseImpl) Delete(id int64) (model.Role, error) {
	return u.RoleRepo.Delete(id)
}

// FindAll implements RoleUsecase
func (u *RoleUsecaseImpl) FindAll() ([]model.Role, error) {
	return u.RoleRepo.FindAll()
}

// FindById implements RoleUsecase
func (u *RoleUsecaseImpl) FindById(id int64) (model.Role, error) {
	return u.RoleRepo.FindById(id)
}

// FindByRolename implements RoleUsecase
func (u *RoleUsecaseImpl) FindByName(name string) (model.Role, error) {
	return u.RoleRepo.FindByName(name)
}

// Save implements RoleUsecase
func (u *RoleUsecaseImpl) Save(newRole model.Role) (model.Role, error) {
	return u.RoleRepo.Save(newRole)
}

// Update implements RoleUsecase
func (u *RoleUsecaseImpl) Update(updatedRole model.Role) (model.Role, error) {
	// Mendapatkan entitas Role sebelumnya dari RoleRepo berdasarkan ID
	previousRole, err := u.RoleRepo.FindById(updatedRole.ID)
	if err != nil {
		return model.Role{}, err
	}

	// Mengambil nilai created_at dari entitas sebelumnya
	createdAt := previousRole.CreatedAt

	// Membuat updatedRole dengan nilai created_at yang ada sebelumnya
	updatedRole.CreatedAt = createdAt

	return u.RoleRepo.Update(updatedRole)
}

func NewRoleUsecaseImpl(RoleRepo repository.RoleRepo) RoleUsecase {
	return &RoleUsecaseImpl{
		RoleRepo: RoleRepo,
	}
}
