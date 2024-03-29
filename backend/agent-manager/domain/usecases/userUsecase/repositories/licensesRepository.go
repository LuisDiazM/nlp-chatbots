package repositories

import "github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"

type LicensesRepoGateway interface {
	GetLicensesByUser(userId string) *[]entities.License
	CreateLicenseByUser(userId string, licenseType string) error
	GetLastLicensesByUser(userId string) *entities.License
	GetLastLicensesUsage(licenseId string) *entities.LicensesUsage
}
