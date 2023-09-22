package usecases

import (
	"context"
	"log"

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

func (usecase *LicenseUsecase) IncrementLicenseUsage(userId string, feature entities.Feature, ctx context.Context) {
	license := usecase.repository.GetLastLicenseByUserId(userId, &ctx)
	if license != nil {
		err := usecase.repository.IncrementLicenseUsageByFeature(*license, feature, &ctx)
		if err != nil {
			log.Printf(`userId :%s cannot update the usage by error %s`, userId, err.Error())
		}
	}
}
