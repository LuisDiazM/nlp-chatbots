package userRepository

import "time"

const (
	queryGetLicenses   = "license.get"
	eventCreateLicense = "license.create"
	timeout            = time.Minute * 1
)
