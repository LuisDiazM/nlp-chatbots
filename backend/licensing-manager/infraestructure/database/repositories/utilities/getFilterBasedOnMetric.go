package utilities

import (
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFilterToUpdateMetricBasedOnFeature(feature entities.Feature) primitive.M {
	switch feature {
	case entities.RateLimit:
		return bson.M{
			"$inc": bson.M{
				"monthlyHistory.$.rateLimit": 1,
			},
		}
	default:
		return bson.M{
			"$inc": bson.M{
				"monthlyHistory.$.trainings": 1,
			},
		}
	}
}
