package entities

type HistoricalRecords struct {
	RateLimit uint  `json:"rate_limit" bson:"rateLimit"`
	Trainings uint  `json:"trainings" bson:"trainings"`
	Month     uint8 `json:"month" bson:"month"`
}

type LicensesUsage struct {
	Year           uint16              `json:"year" bson:"year"`
	LicenseId      string              `json:"license_id" bson:"licenseId"`
	MonthlyHistory []HistoricalRecords `json:"monthly_history" bson:"monthlyHistory"`
	Id             string              `json:"id" bson:"_id,omitempty"`
}
