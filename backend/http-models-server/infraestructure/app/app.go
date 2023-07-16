package app

import (
	"context"
	"fmt"
	"http-models-server/cmd/config"
	trainingusecase "http-models-server/domain/usecases/trainingUsecase"
	"http-models-server/infraestructure/database"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	WebServer       *gin.Engine
	Env             *config.Env
	Database        *database.DatabaseImp
	TrainingUsecase *trainingusecase.TrainingUsecase
	// ModelsUsecase   *modelusecase.ModelsUsecase
}

func NewApplication(webServer *gin.Engine, configVars *config.Env, database *database.DatabaseImp, trainingUsecase *trainingusecase.TrainingUsecase) *Application {
	return &Application{WebServer: webServer,
		Env:             configVars,
		Database:        database,
		TrainingUsecase: trainingUsecase,
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
	return g.Wait()
}
