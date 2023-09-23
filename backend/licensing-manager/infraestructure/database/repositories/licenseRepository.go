package repositories

import (
	"context"
	"log"
	"time"

	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/entities"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/repositories"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/database"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/database/repositories/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repository *LicenseRepository) GetLastLicenseByUserId(userId string, ctx *context.Context) *entities.License {
	collection := repository.Db.Collection(databaseName, collectionLicenses)
	filter := primitive.M{"userId": userId}
	optionsSort := options.FindOne().SetSort(bson.M{"expiredAt": -1})
	result := collection.FindOne(*ctx, filter, optionsSort)
	var license entities.License
	err := result.Decode(&license)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &license
}

func (repository *LicenseRepository) IncrementLicenseUsageByFeature(license entities.License, feature entities.Feature, ctx *context.Context) error {
	collection := repository.Db.Collection(databaseName, collectionLicensesUsage)
	optionsSort := options.FindOne().SetSort(bson.M{"year": -1})
	filter := primitive.M{"licenseId": license.Id}
	result := collection.FindOne(*ctx, filter, optionsSort)
	var licenseUsage entities.LicensesUsage
	var licenseUsageId primitive.ObjectID
	err := result.Decode(&licenseUsage)
	timestamp := time.Now()
	year := timestamp.Year()
	month := timestamp.Month()
	if err != nil {
		log.Println(err)
		licenseUsage := entities.LicensesUsage{Year: uint16(year), LicenseId: license.Id, MonthlyHistory: []entities.HistoricalRecords{{Month: uint8(month)}}}
		result := repository.Db.InsertOne(collection, ctx, licenseUsage)
		licenseUsageId = result.InsertedID.(primitive.ObjectID)
	} else {
		licenseUsageId, err = primitive.ObjectIDFromHex(licenseUsage.Id)
		if err != nil {
			log.Println(err)
		}
	}
	filterUsage := bson.M{"_id": licenseUsageId, "monthlyHistory.month": month}

	updateLicenseUsage := utilities.GetFilterToUpdateMetricBasedOnFeature(feature)
	_, err = collection.UpdateOne(*ctx, filterUsage, updateLicenseUsage)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repository *LicenseRepository) GetLastLicenseUsageById(licenseId string, ctx *context.Context) *entities.LicensesUsage {
	collection := repository.Db.Collection(databaseName, collectionLicensesUsage)
	filter := primitive.M{"licenseId": licenseId}
	optionsSort := options.FindOne().SetSort(bson.M{"year": -1})
	result := collection.FindOne(*ctx, filter, optionsSort)
	var license entities.LicensesUsage
	err := result.Decode(&license)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &license
}
