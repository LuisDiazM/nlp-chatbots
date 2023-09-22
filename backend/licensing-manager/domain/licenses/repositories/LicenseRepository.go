package repositories

import (
	"context"

	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/entities"
)

type LicenseRepositoryGateway interface {
	CreateLicense(license entities.License, ctx *context.Context) *interface{}
	FindLicensesByUserId(userId string, ctx *context.Context) *[]entities.License
	GetLastLicenseByUserId(userId string, ctx *context.Context) *entities.License
	IncrementLicenseUsageByFeature(license entities.License, feature entities.Feature, ctx *context.Context) error
}
