package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/module/auth"
	"github.com/root9464/Ton-students/backend/shared/logger"
)

type App struct {
	app        *fiber.App
	auhtModule *auth.AuthModule

	Config *config.Config
	logger *logger.Logger
}

func NewApp(config *config.Config, logger *logger.Logger) (*App, error) {
	a := &App{
		app:    fiber.New(),
		Config: config,
		logger: logger,
	}

	return a, nil
}

func (a *App) Run() error {
	a.app.Use(config.CORS_CONFIG)
	err := a.initDep()
	if err != nil {
		return err
	}
	return a.runHttpServer()
}

func (a *App) initDep() error {
	inits := []func() error{
		a.initConfig,
		a.initAuthModule,
		a.initRouter,
	}

	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig() error {
	err := config.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initRouter() error {
	api := a.app.Group("/api")
	a.auhtModule.Rotes(api)
	return nil
}

func (a *App) initAuthModule() error {
	a.auhtModule = auth.NewAuthModule(a.Config, a.logger)
	if a.auhtModule == nil {
		return fmt.Errorf("failed to create auth module")
	}
	return nil
}

func (a *App) runHttpServer() error {
	err := a.app.Listen(fmt.Sprintf(":%v", 6069))
	if err != nil {
		return err
	}
	return nil
}
