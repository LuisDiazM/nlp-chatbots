package entities

import "time"

type LicenseFeature struct {
	RateLimit int `json:"rate_limit,omitempty"`
	Trainings int `json:"trainings,omitempty"`
}

type License struct {
	Id        string         `json:"id,omitempty" `
	ExpiredAt time.Time      `json:"expired_at,omitempty" `
	UserId    string         `json:"user_id,omitempty" `
	CreatedAt time.Time      `json:"created_at,omitempty"`
	Features  LicenseFeature `json:"features,omitempty" `
	Type      string         `json:"type,omitempty"`
}

const (
	LICENSE_FREE = "FREE"
)
