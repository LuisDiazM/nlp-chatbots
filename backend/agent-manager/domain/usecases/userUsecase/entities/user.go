package entities

import "time"

type User struct {
	Name      string    `json:"name,omitempty" bson:"name"`
	Email     string    `json:"email,omitempty" bson:"email"`
	Id        string    `json:"id,omitempty" bson:"_id"`
	Picture   string    `json:"picture,omitempty" bson:"picture"`
	CreatedAt time.Time `json:"created_at" bson:"createdAt"`
}

type UserWithLicenseValidation struct {
	User           User `json:"user"`
	IsLicenceValid bool `json:"is_licence_valid"`
}
