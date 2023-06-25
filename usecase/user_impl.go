package usecase

import (
	"time"

	"github.com/sferawann/pinjol/model"
	"github.com/sferawann/pinjol/repository"
	"github.com/sferawann/pinjol/utils"
)

type UserUsecaseImpl struct {
	UserRepo repository.UserRepo
	RoleRepo repository.RoleRepo
}

// Delete implements UserUsecase
func (u *UserUsecaseImpl) Delete(id int64) (model.User, error) {
	return u.UserRepo.Delete(id)
}

// FindAll implements UserUsecase
func (u *UserUsecaseImpl) FindAll() ([]model.User, error) {
	return u.UserRepo.FindAll()
}

// FindById implements UserUsecase
func (u *UserUsecaseImpl) FindById(id int64) (model.User, error) {
	return u.UserRepo.FindById(id)
}

// FindByUsername implements UserUsecase
func (u *UserUsecaseImpl) FindByUsername(username string) (model.User, error) {
	return u.UserRepo.FindByUsername(username)
}

// Save implements UserUsecase
func (u *UserUsecaseImpl) Save(newUser model.User) (model.User, error) {
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return model.User{}, err
	}

	newUser.Password = hashedPassword
	newUser.Limit = 2000000
	newUser.RoleID = 1

	role, err := u.RoleRepo.FindById(newUser.RoleID)
	if err != nil {
		return model.User{}, err
	}

	newUser.Role = role

	return u.UserRepo.Save(newUser)
}

// Update implements UserUsecase
func (u *UserUsecaseImpl) Update(updatedUser model.User) (model.User, error) {
	hashedPassword, err := utils.HashPassword(updatedUser.Password)
	if err != nil {
		return model.User{}, err
	}

	// Mendapatkan entitas User sebelumnya dari UserRepo berdasarkan ID
	previousUser, err := u.UserRepo.FindById(updatedUser.ID)
	if err != nil {
		return model.User{}, err
	}

	// Mengambil nilai-nilai field dari entitas sebelumnya
	previousUsername := previousUser.Username
	previousPassword := previousUser.Password
	previousKTP := previousUser.KTP
	previousName := previousUser.Name
	previousAddress := previousUser.Address
	previousPhoneNumber := previousUser.PhoneNumber
	previousLimit := previousUser.Limit
	previousRoleID := previousUser.RoleID
	previousCreatedAt := previousUser.CreatedAt

	// Menggunakan nilai-nilai field sebelumnya untuk field-field yang tidak diubah
	if updatedUser.Username == "" {
		updatedUser.Username = previousUsername
	}
	if updatedUser.Password == "" {
		updatedUser.Password = previousPassword
	}
	if updatedUser.KTP == "" {
		updatedUser.KTP = previousKTP
	}
	if updatedUser.Name == "" {
		updatedUser.Name = previousName
	}
	if updatedUser.Address == "" {
		updatedUser.Address = previousAddress
	}
	if updatedUser.PhoneNumber == "" {
		updatedUser.PhoneNumber = previousPhoneNumber
	}
	if updatedUser.Limit == 0 {
		updatedUser.Limit = previousLimit
	}
	if updatedUser.RoleID == 0 {
		updatedUser.RoleID = previousRoleID
	}
	if updatedUser.CreatedAt == (time.Time{}) {
		updatedUser.CreatedAt = previousCreatedAt
	}

	// Hash password baru jika ada perubahan
	if updatedUser.Password != previousPassword {
		updatedUser.Password = hashedPassword
	}

	return u.UserRepo.Update(updatedUser)
}

func NewUserUsecaseImpl(UserRepo repository.UserRepo, RoleRepo repository.RoleRepo) UserUsecase {
	return &UserUsecaseImpl{
		UserRepo: UserRepo,
		RoleRepo: RoleRepo,
	}
}
