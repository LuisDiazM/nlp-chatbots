package userusecase

import (
	"context"
	"time"

	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/repositories"
	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/utilities"
)

type UserUsecase struct {
	DatabaseRepository  repositories.UserRepositoryGateway
	MessagingRepository repositories.LicensesRepoGateway
}

func NewUserUsecase(databaseRepo repositories.UserRepositoryGateway, messagingRepo repositories.LicensesRepoGateway) *UserUsecase {
	return &UserUsecase{DatabaseRepository: databaseRepo, MessagingRepository: messagingRepo}
}

func (usecase *UserUsecase) GetUserById(id string, ctx context.Context) *entities.UserWithLicenseValidation {
	user := usecase.DatabaseRepository.GetUserById(id, ctx)
	var userWithLicenseValidation entities.UserWithLicenseValidation
	if user == nil {
		return nil
	} else {
		licenses := usecase.MessagingRepository.GetLicensesByUser(user.Email)
		timestamp := time.Now()
		isLicenseValid := utilities.CheckAvalibleLicensesByDate(timestamp, *licenses)
		userWithLicenseValidation = entities.UserWithLicenseValidation{User: *user, IsLicenceValid: isLicenseValid}
	}
	return &userWithLicenseValidation
}

func (usecase *UserUsecase) InsertUser(user entities.User, ctx context.Context) *map[string]bool {
	userData := usecase.GetUserById(user.Id, ctx)
	if userData == nil {
		licenses := usecase.MessagingRepository.GetLicensesByUser(user.Id)
		if len(*licenses) > 0 {
			var licenseAnalisis map[string]bool = map[string]bool{"exists_previous_licenses": true}
			return &licenseAnalisis
		} else {
			response := usecase.DatabaseRepository.InsertUser(user, ctx)
			if response == nil {
				return nil
			}
			err := usecase.MessagingRepository.CreateLicenseByUser(user.Id, entities.LICENSE_FREE)
			var licenseAnalisis map[string]bool = map[string]bool{"exists_previous_licenses": false}
			if err != nil {
				return nil
			}
			return &licenseAnalisis
		}
	}
	return nil
}
