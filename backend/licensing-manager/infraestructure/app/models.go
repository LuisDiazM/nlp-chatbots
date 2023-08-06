package app

type RequestCreateLicense struct {
	UserId      string `json:"user_id,omitempty"`
	LicenseType string `json:"license_type,omitempty"`
}

type RequestLicense struct {
	UserId string `json:"user_id,omitempty"`
}
