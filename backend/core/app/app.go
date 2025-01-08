package app

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
	"github.com/root9464/Ton-students/config"
	"github.com/root9464/Ton-students/database"
	redis_connect "github.com/root9464/Ton-students/redis"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/root9464/Ton-students/shared/middleware"
	"gorm.io/gorm"
)

type App struct {
	app *fiber.App

	logger         *logger.Logger
	db             *gorm.DB
	redis          *redis.Client
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

func (app *App) Run() error {
	app.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://0.0.0.0:5173, https://4f67-95-105-125-55.ngrok-free.app",
		AllowCredentials: true,
	}))
	app.app.Use(middleware.LoggerMiddleware())

	app.initDeps()

	return app.runHttpServer()
}

func (app *App) initDeps() {
	inits := []func() error{
		app.initConfig,
		app.initDb,
		app.initRedis,
		app.initLogger,
		app.initValidator,
		app.initModuleProvider,
		app.initRouter,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			log.Fatalf("✖ Failed to initialize dependencies: %s", err.Error())
		}
	}
}

func (app *App) initConfig() error {
	if app.config == nil {
		config, err := config.LoadConfig(".")
		if err != nil {
			return fmt.Errorf("✖ Failed to load config: %s", err.Error())
		}
		app.config = &config
	}

	err := config.Load(".env")
	if err != nil {
		return fmt.Errorf("✖ Failed to load config: %s", err.Error())
	}

	return nil
}

func (app *App) initDb() error {
	if app.db == nil {
		db, err := database.ConnectDb(app.config.DatabaseUrl)
		if err != nil {
			app.logger.Errorf("✖ Failed to connect to database: %s", err.Error())
			return err
		}
		app.db = db
		// true - миграция
		// false - не миграция
		if err := database.Migrate(db, false); err != nil {
			return fmt.Errorf("✖ Failed to migrate database: %s", err.Error())
		}
	}

	return nil
}

func (app *App) initRedis() error {
	if app.redis == nil {
		redis, err := redis_connect.Connect(app.config.RedisUrl)
		if err != nil {
			return fmt.Errorf("✖ Failed to connect to redis: %s", err.Error())
		}
		app.redis = redis

		// 0 - не трогать кэш
		// 1 - выборочная очистка
		// 2 - полная очистка
		if err := redis_connect.FlushRedisCache(redis, 0); err != nil {
			return fmt.Errorf("✖ Failed to flush redis cache: %s", err.Error())
		}
	}
	return nil
}

func (app *App) initLogger() error {
	if app.logger == nil {
		app.logger = logger.GetLogger()
	}
	return nil
}

func (app *App) initValidator() error {
	if app.validator == nil {
		app.validator = validator.New()
	}
	return nil
}

func (app *App) initModuleProvider() error {
	var err error
	app.moduleProvider, err = NewModuleProvider(app)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) runHttpServer() error {
	if app.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			return fmt.Errorf("✖ Failed to load config: %s", err.Error())
		}
		app.httpConfig = cfg
	}

	log.Infof("🌐 Server is running on %s", app.httpConfig.Address())
	log.Info("✅ Server started successfully")
	if err := app.app.Listen(app.httpConfig.Address()); err != nil {
		return fmt.Errorf("✖ Failed to start http server: %s", err.Error())
	}

	return nil
}

func (app *App) initRouter() error {
	api := app.app.Group("/api")

	app.moduleProvider.userModule.UserRoutes(api)
	app.moduleProvider.authModule.AuthRoutes(api)

	creator := api.Group("/creator")
	app.moduleProvider.serviceModule.ServiceRoutes(creator)

	return nil
}
