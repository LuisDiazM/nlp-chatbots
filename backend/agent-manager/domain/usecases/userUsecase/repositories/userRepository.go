package repositories

import (
	"context"

	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"
)

type UserRepositoryGateway interface {
	InsertUser(user entities.User, ctx context.Context) *interface{}
	GetUserById(id string, ctx context.Context) *entities.User
	DeleteUserById(id string, ctx context.Context) error
}
