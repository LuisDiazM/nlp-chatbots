package app

import (
	"context"
	"fmt"
	"log"

	"github.com/LuisDiazM/agent-manager/cmd/config"
	nnmodelusecase "github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase"
	trainingusecase "github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase"
	userusecase "github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase"

	"github.com/LuisDiazM/agent-manager/infraestructure/database"
	"github.com/LuisDiazM/agent-manager/infraestructure/messaging"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	WebServer            *gin.Engine
	Env                  *config.Env
	Database             *database.DatabaseImp
	Messaging            *messaging.NatsImp
	TrainingUsecase      *trainingusecase.TrainingUsecase
	UserUsecase          *userusecase.UserUsecase
	NeuralNetworkUsecase *nnmodelusecase.NeuralNetworkModelUsecase
}

func NewApplication(webServer *gin.Engine, configVars *config.Env, database *database.DatabaseImp, trainingUsecase *trainingusecase.TrainingUsecase, userUsecase *userusecase.UserUsecase, messaging *messaging.NatsImp, nnmodelUsecase *nnmodelusecase.NeuralNetworkModelUsecase) *Application {
	return &Application{WebServer: webServer,
		Env:                  configVars,
		Database:             database,
		TrainingUsecase:      trainingUsecase,
		UserUsecase:          userUsecase,
		Messaging:            messaging,
		NeuralNetworkUsecase: nnmodelUsecase,
	}
}

func (app *Application) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := app.WebServer.Run(fmt.Sprintf(`:%d`, app.Env.PORT)); err != nil {
			return err
		}
		return nil
	})
	log.Println("agent manager started ...")

	return g.Wait()
}
