package app

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

	config     *config.Config
	logger     *logger.Logger
	validator  *validator.Validate
	httpConfig config.HTTPConfig

	db    *gorm.DB
	redis *redis.Client

	moduleProvider *moduleProvider
}

func NewApp() *App {
	return &App{
		app: fiber.New(),
	}
}

var wg sync.WaitGroup

func (app *App) Run() error {
	app.app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
	}))
	app.app.Use(middleware.LoggerMiddleware())

	err := app.initDeps()

	if err != nil {
		return err
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := app.runHttpServer(); err != nil {
			app.logger.Errorf("%s", "✖ Failed to start server: "+err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if err := app.initBot(); err != nil {
			app.logger.Errorf("%s", "✖ Failed to start bot: "+err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func (app *App) initDeps() error {

	inits := []func() error{
		app.initConfig,
		app.initLogger,
		app.initValidator,

		app.initDb,
		app.initRedis,

		app.initModuleProvider,
		app.initRouter,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			return fmt.Errorf("%s", "✖ Failed to initialize dependencies: "+err.Error())
		}
	}
	return nil
}

func (app *App) initConfig() error {
	if app.config == nil {
		config, err := config.LoadConfig(".")
		if err != nil {
			return fmt.Errorf("%s", "✖ Failed to load config: "+err.Error())
		}
		app.config = config
	}

	err := config.Load("../.env")
	if err != nil {
		return fmt.Errorf("%s", "✖ Failed to load config: "+err.Error())
	}

	return nil
}

func (app *App) initDb() error {
	if app.db == nil {
		db, err := database.ConnectDb(app.config.DatabaseUrl, app.logger)
		if err != nil {
			return err
		}
		app.db = db

		// true - запустить миграцию
		// false - не запускать
		if err := database.Migrate(db, false, app.logger); err != nil {
			return fmt.Errorf("%s", "✖ Failed to migrate database: "+err.Error())
		}
	}

	return nil
}

func (app *App) initRedis() error {
	if app.redis == nil {
		redis, err := redis_connect.Connect(app.config.RedisUrl, app.logger)
		if err != nil {
			app.logger.Errorf("Failed to connect to Redis: %v", err)
			return nil
		}
		app.redis = redis

		// 0 - не трогать кэш
		// 1 - выборочная очистка
		// 2 - полная очистка
		if err := redis_connect.FlushRedisCache(redis, 0, app.logger); err != nil {
			err = fmt.Errorf("✖ Failed to flush redis cache: %v", err)
			app.logger.Errorf("%s", err.Error())
			return err
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
	err := error(nil)
	app.moduleProvider, err = NewModuleProvider(app)
	if err != nil {
		app.logger.Errorf("%s", err.Error())
		return err
	}
	return nil
}

func (app *App) runHttpServer() error {
	if app.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			app.logger.Errorf("%s", "✖ Failed to load config: "+err.Error())
			return fmt.Errorf("✖ Failed to load config: %v", err)
		}
		app.httpConfig = cfg
	}

	app.logger.Infof("🌐 Server is running on %s", app.httpConfig.Address())
	app.logger.Info("✅ Server started successfully")
	if err := app.app.Listen(app.httpConfig.Address()); err != nil {
		app.logger.Errorf("%s", "✖ Failed to start server: "+err.Error())
		return fmt.Errorf("✖ Failed to start server: %v", err)
	}

	return nil
}

func (app *App) initBot() error {
	if app.moduleProvider.botModule == nil {
		app.logger.Warnf("Bot module is not initialized")
		return nil
	}

	if err := app.moduleProvider.botModule.InitBot(); err != nil {
		app.logger.Errorf("%s", "✖ Failed to start bot: "+err.Error())
	}

	return nil
}

func (app *App) initRouter() error {
	api := app.app.Group("/api")

	app.moduleProvider.authModule.AuthRoutes(api)
	app.moduleProvider.userModule.UserRoutes(api)

	app.moduleProvider.serviceModule.ServiceRoutes(api)

	app.moduleProvider.botModule.BotRoutes(api)

	// app.moduleProvider.notificationsModule.NotificationsRoutes(api)

	return nil
}
