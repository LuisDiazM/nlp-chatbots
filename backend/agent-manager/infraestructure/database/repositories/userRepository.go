package repositories

import (
	"context"
	"log"

	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/repositories"
	"github.com/LuisDiazM/agent-manager/infraestructure/database"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	Database *database.DatabaseImp
}

func NewUserRepository(db *database.DatabaseImp) repositories.UserRepositoryGateway {
	return &UserRepository{
		Database: db,
	}
}

func (userRepository *UserRepository) InsertUser(user entities.User, ctx context.Context) *interface{} {
	collection := userRepository.Database.Collection(trainingDatabaseName, userCollectionName)
	result := userRepository.Database.InsertOne(collection, &ctx, user)
	if result != nil {
		return &result.InsertedID
	} else {
		return nil
	}
}

func (userRepository *UserRepository) GetUserById(id string, ctx context.Context) *entities.User {
	collection := userRepository.Database.Collection(trainingDatabaseName, userCollectionName)
	var userInfo entities.User
	filter := bson.M{"_id": id}
	data := collection.FindOne(ctx, filter)
	err := data.Decode(&userInfo)
	if err != nil {
		log.Println(id, err)
		return nil
	}
	return &userInfo
}

func (repository *UserRepository) DeleteUserById(id string, ctx context.Context) error {
	collection := repository.Database.Collection(trainingDatabaseName, userCollectionName)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("mongo -> DeleteOne", err)
		return nil
	}
	return nil
}
