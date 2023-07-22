package userusecase

import (
	"context"

	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/repositories"
)

type UserUsecase struct {
	DatabaseRepository repositories.UserRepositoryGateway
}

func NewUserUsecase(databaseRepo repositories.UserRepositoryGateway) *UserUsecase {
	return &UserUsecase{DatabaseRepository: databaseRepo}
}

func (usecase *UserUsecase) GetUserById(id string, ctx context.Context) *entities.UserWithLicenseValidation {
	// debo validar la licencia
	user := usecase.DatabaseRepository.GetUserById(id, ctx)
	var userWithLicenseValidation entities.UserWithLicenseValidation
	if user == nil {
		return nil
	} else {
		userWithLicenseValidation = entities.UserWithLicenseValidation{User: *user, IsLicenceValid: true}
	}
	return &userWithLicenseValidation
}

func (usecase *UserUsecase) InsertUser(user entities.User, ctx context.Context) *interface{} {
	// debo emitir el evento de que se creó un usuario
	return usecase.DatabaseRepository.InsertUser(user, ctx)
}
