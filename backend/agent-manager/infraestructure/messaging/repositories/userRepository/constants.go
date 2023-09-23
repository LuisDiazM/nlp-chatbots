package userRepository

import "time"

const (
	queryGetLicenses     = "license.get"
	queryGetLastLicense  = "license.get.last.byUserId"
	queryGetLicenseUsage = "license.usage.get"
	eventCreateLicense   = "license.create"
	timeout              = time.Minute * 1
)
