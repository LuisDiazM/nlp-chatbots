package models

import "time"

type ChannelInfo struct {
	ChannelId   string    `bson:"channelId,omitempty" json:"channel_id,omitempty"`
	ChannelType string    `bson:"channelType,omitempty" json:"channel_type,omitempty"`
	ModelId     string    `bson:"modelId,omitempty" json:"model_id,omitempty"`
	Created     time.Time `bson:"created,omitempty" json:"created,omitempty"`
	State       string    `bson:"state,omitempty" json:"state,omitempty"`
	Id          string    `bson:"_id" json:"id,omitempty"`
	UserId      string    `bson:"userId" json:"user_id,omitempty"`
}

type TrainingInfo struct {
	Intents []Intent `bson:"intents" json:"intents,omitempty"`
	UserId  string   `bson:"userId" json:"user_id,omitempty"`
	Id      string   `bson:"_id,omitempty" json:"id,omitempty"`
}

type Intent struct {
	Patterns  []string `bson:"patterns" json:"patterns,omitempty"`
	Responses []string `bson:"responses" json:"responses"`
	Tag       string   `bson:"tag" json:"tag,omitempty"`
}
