package entities

import "time"

type NNModel struct {
	ModelName   string    `json:"model_name,omitempty" bson:"modelName,omitempty"`
	BucketName  string    `json:"bucket_name,omitempty" bson:"bucketName,omitempty"`
	NluIntentId string    `json:"nlu_intent_id,omitempty" bson:"nluIntentId,omitempty"`
	Created     time.Time `json:"created,omitempty" bson:"created,omitempty"`
	UserId      string    `json:"user_id,omitempty" bson:"userId,omitempty"`
	Id          string    `json:"id,omitempty" bson:"_id,omitempty"`
}
