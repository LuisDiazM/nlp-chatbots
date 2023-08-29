package entities

type Intent struct {
	Patterns  []string `json:"patterns,omitempty" bson:"patterns"`
	Responses []string `json:"responses,omitempty" bson:"responses"`
	Tag       string   `json:"tag,omitempty" bson:"tag"`
}

type TrainingInfo struct {
	UserId      string   `json:"user_id,omitempty" bson:"userId" binding:"required"`
	Intents     []Intent `json:"intents,omitempty" bson:"intents" binding:"required"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Id          string   `json:"id,omitempty" bson:"_id,omitempty"`
}
