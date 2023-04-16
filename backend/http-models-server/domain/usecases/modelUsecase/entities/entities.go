package entities

import "time"

type Model struct {
	ModelName      string    `json:"model_name,omitempty" bson:"modelName"`
	BucketName     string    `json:"bucket_name,omitempty" bson:"bucketName"`
	TrainingDataId string    `json:"training_data_id,omitempty" bson:"trainingDataId"`
	Created        time.Time `json:"created,omitempty" bson:"created"`
}
