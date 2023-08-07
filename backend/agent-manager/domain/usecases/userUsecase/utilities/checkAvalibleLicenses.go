package utilities

import (
	"time"

	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"
)

func CheckAvalibleLicensesByDate(dateTime time.Time, licenses []entities.License) bool {
	if len(licenses) == 0 {
		return false
	}
	for _, license := range licenses {
		if license.ExpiredAt.Unix() > dateTime.Unix() {
			return true
		}
	}
	return false
}
