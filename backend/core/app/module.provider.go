package app

import (
	auth_module "github.com/root9464/Ton-students/module/auth"
	bot_module "github.com/root9464/Ton-students/module/bot"
	chat_module "github.com/root9464/Ton-students/module/chat"
	jwt_module "github.com/root9464/Ton-students/module/jwt"
	notifications_module "github.com/root9464/Ton-students/module/notifications"
	service_module "github.com/root9464/Ton-students/module/service_module"
	user_module "github.com/root9464/Ton-students/module/user"
)

type moduleProvider struct {
	userModule          *user_module.UserModule
	authModule          *auth_module.AuthModule
	serviceModule       *service_module.ServiceModule
	botModule           *bot_module.BotModule
	chatModule          *chat_module.ChatModule
	notificationsModule *notifications_module.NotificationsModule

	jwtModule *jwt_module.JwtModule
	app       *App
}

func NewModuleProvider(app *App) (*moduleProvider, error) {
	provider := &moduleProvider{
		app: app,
	}

	err := provider.initDeps()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *moduleProvider) initDeps() error {
	inits := []func() error{
		p.JwtModule,
		p.UserModule,
		p.AuthModule,
		p.ServiceModule,
		p.BotModule,
		p.NotificationsModule,
		p.ChatModule,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			p.app.logger.Errorf("%s", "✖ Failed to initialize module: "+err.Error())
			return err
		}
	}
	return nil
}

func (p *moduleProvider) UserModule() error {
	p.userModule = user_module.NewUserModule(p.app.logger, p.app.validator, p.app.db, p.app.config.JwtPublicKey, p.app.config.HmacKey, *p.jwtModule)
	return nil
}

func (p *moduleProvider) JwtModule() error {
	p.jwtModule = jwt_module.NewJwtModule(p.app.logger, p.app.validator, p.app.db, p.app.config.JwtPrivateKey, p.app.config.JwtPublicKey)
	return nil
}

func (p *moduleProvider) AuthModule() error {
	p.authModule = auth_module.NewAuthModule(p.app.logger, p.app.validator, p.app.config, p.userModule.UserService(), *p.jwtModule)
	return nil
}

func (p *moduleProvider) ServiceModule() error {
	p.serviceModule = service_module.NewServiceModule(p.app.logger, p.app.validator, p.app.db, p.userModule.UserRepo(), *p.jwtModule, p.app.config.JwtPublicKey)
	return nil
}

func (p *moduleProvider) ChatModule() error {
	p.chatModule = chat_module.NewChatModule()
	return nil
}

func (p *moduleProvider) BotModule() error {
	p.botModule = bot_module.NewBotModule(p.app.logger, p.app.validator, p.app.db, p.app.config, *p.userModule, *p.jwtModule)
	return nil
}

func (p *moduleProvider) NotificationsModule() error {
	p.notificationsModule = notifications_module.NewNotificationsModule(p.app.logger, p.app.validator, p.app.db, *p.userModule)
	return nil
}
