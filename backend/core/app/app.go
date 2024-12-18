package app

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/root9464/Ton-students/config"
	"github.com/root9464/Ton-students/ent"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/root9464/Ton-students/shared/middleware"
)

type App struct {
	app *fiber.App

	logger         *logger.Logger
	db             *ent.Client
	validator      *validator.Validate
	config         *config.Config
	httpConfig     config.HTTPConfig
	moduleProvider *moduleProvider
}

func NewApp() *App {
	return &App{
		app: fiber.New(),
	}
}

func (a *App) Run() error {
	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://0.0.0.0:5173, https://4f67-95-105-125-55.ngrok-free.app",
		AllowCredentials: true,
	}))
	a.app.Use(middleware.LoggerMiddleware())

	a.initDeps()

	return a.runHttpServer()
}

func (a *App) initDeps() {
	inits := []func() error{
		a.initConfig,
		a.initDb,
		a.initLogger,
		a.initValidator,
		a.initModuleProvider,
		a.initRouter,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			log.Fatalf("✖ Failed to initialize dependencies: %s", err.Error())
		}
	}
}

func (a *App) initConfig() error {
	if a.config == nil {
		config, err := config.LoadConfig(".")
		if err != nil {
			return fmt.Errorf("✖ Failed to load config: %s", err.Error())
		}
		a.config = &config
	}

	err := config.Load(".env")
	if err != nil {
		return fmt.Errorf("✖ Failed to load config: %s", err.Error())
	}

	return nil
}

func (a *App) initDb() error {
	if a.db == nil {
		db, err := ent.Open("postgres", a.config.DatabaseUrl)
		if err != nil {
			return fmt.Errorf("✖ Failed to connect to database: %s", err.Error())
		}
		a.db = db

		if err := db.Schema.Create(context.Background()); err != nil {
			return fmt.Errorf("✖ Failed to create schema resources: %s", err.Error())
		}
	}

	return nil
}

func (a *App) initLogger() error {
	if a.logger == nil {
		a.logger = logger.GetLogger()
	}
	return nil
}

func (a *App) initValidator() error {
	if a.validator == nil {
		a.validator = validator.New()
	}
	return nil
}

func (a *App) initModuleProvider() error {
	var err error
	a.moduleProvider, err = NewModuleProvider(a)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) runHttpServer() error {
	if a.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			return fmt.Errorf("✖ Failed to load config: %s", err.Error())
		}
		a.httpConfig = cfg
	}

	log.Infof("🌐 Server is running on %s", a.httpConfig.Address())
	log.Info("✅ Server started successfully")
	if err := a.app.Listen(a.httpConfig.Address()); err != nil {
		return fmt.Errorf("✖ Failed to start http server: %s", err.Error())
	}

	return nil
}

func (a *App) initRouter() error {
	api := a.app.Group("/api")

	a.moduleProvider.authModule.AuthRoutes(api)
	a.moduleProvider.userModule.UserRoutes(api)
	return nil
}
