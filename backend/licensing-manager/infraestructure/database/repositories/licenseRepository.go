package repositories

import (
	"context"
	"log"

	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/entities"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/repositories"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LicenseRepository struct {
	Db *database.DatabaseImp
}

func NewLicenseRepository(db *database.DatabaseImp) repositories.LicenseRepositoryGateway {
	return &LicenseRepository{Db: db}
}

func (repository *LicenseRepository) CreateLicense(license entities.License, ctx *context.Context) *interface{} {
	collection := repository.Db.Collection(databaseName, collectionLicenses)
	result := repository.Db.InsertOne(collection, ctx, license)
	if result != nil {
		return &result.InsertedID
	} else {
		return nil
	}
}

func (repository *LicenseRepository) FindLicensesByUserId(userId string, ctx *context.Context) *[]entities.License {
	collection := repository.Db.Collection(databaseName, collectionLicenses)
	filter := primitive.M{"userId": userId}
	cursor, err := repository.Db.Find(collection, ctx, filter)
	if err != nil {
		log.Println(err)
		return nil
	}
	var licenses []entities.License
	err = cursor.All(*ctx, &licenses)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &licenses
}
