package entities

import (
	"time"
)

type LicenseFeature struct {
	RateLimit int `json:"rate_limit,omitempty" bson:"rateLimit"`
	Trainings int `json:"trainings,omitempty" bson:"trainings"`
}

type License struct {
	Id        string         `json:"id,omitempty" bson:"_id,omitempty"`
	ExpiredAt time.Time      `json:"expired_at,omitempty" bson:"expiredAt"`
	UserId    string         `json:"user_id,omitempty" bson:"userId"`
	CreatedAt time.Time      `json:"created_at,omitempty" bson:"createdAt"`
	Features  LicenseFeature `json:"features,omitempty" bson:"features"`
	Type      string         `json:"type,omitempty" bson:"type"`
}
