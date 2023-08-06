package usecases

import (
	"context"

	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/entities"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/repositories"
)

type LicenseUsecase struct {
	repository repositories.LicenseRepositoryGateway
}

func NewLicenseUseCase(repo repositories.LicenseRepositoryGateway) LicenseUsecase {
	return LicenseUsecase{repository: repo}
}

func (usecase *LicenseUsecase) CreateLicense(license entities.License, ctx context.Context) *interface{} {
	return usecase.repository.CreateLicense(license, &ctx)
}

func (usecase *LicenseUsecase) GetLicensesByUserId(userId string, ctx context.Context) *[]entities.License {
	return usecase.repository.FindLicensesByUserId(userId, &ctx)
}
