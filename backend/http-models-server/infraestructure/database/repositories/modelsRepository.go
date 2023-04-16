package repositories

import (
	"context"
	"http-models-server/domain/usecases/modelUsecase/entities"
	"http-models-server/domain/usecases/modelUsecase/repositories"
	"http-models-server/infraestructure/database"
	"log"
)

type ModelsRepository struct {
	Database database.DatabaseImp
}

func NewModelRepository(db database.DatabaseImp) repositories.ModelsRepository {
	return &ModelsRepository{
		Database: db,
	}
}

func (repository *ModelsRepository) GetModelNLP(id string, ctx context.Context) *entities.Model {
	collection := repository.Database.Collection(trainingDatabaseName, modelsCollectionName)
	var model entities.Model
	data := repository.Database.FindOne(collection, &ctx, id)
	err := data.Decode(&model)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &model
}
