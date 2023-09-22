package app

import "github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/entities"

type RequestCreateLicense struct {
	UserId      string `json:"user_id,omitempty"`
	LicenseType string `json:"license_type,omitempty"`
}

type RequestLicense struct {
	UserId string `json:"user_id,omitempty"`
}

type RequestIncrementLicenseUsage struct {
	UserId  string           `json:"user_id,omitempty"`
	Feature entities.Feature `json:"feature,omitempty"`
}
