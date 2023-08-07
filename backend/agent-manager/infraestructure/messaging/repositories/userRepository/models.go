package userRepository

type RequestGetLicenses struct {
	UserId string `json:"user_id,omitempty"`
}

type RequestCreateLicense struct {
	UserId      string `json:"user_id,omitempty"`
	LicenseType string `json:"license_type,omitempty"`
}
