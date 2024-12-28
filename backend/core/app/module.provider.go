package app

import (
	auth_module "github.com/root9464/Ton-students/module/auth"
	user_module "github.com/root9464/Ton-students/module/user"
	bot_model "github.com/root9464/Ton-students/module/bot"
	chat_module "github.com/root9464/Ton-students/module/chat"
)

type moduleProvider struct {
	authModule *auth_module.AuthModule
	userModule *user_module.UserModule
	botModule  *bot_model.BotModule
	chatModule *chat_module.ChatModule

	app *App
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
		p.UserModule,
		p.AuthModule,
		p.BotModule,
		p.ChatModule,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *moduleProvider) UserModule() error {
	p.userModule = user_module.NewUserModule(p.app.logger, p.app.validator)
	return nil
}

func (p *moduleProvider) AuthModule() error {
	p.authModule = auth_module.NewAuthModule(p.app.logger, p.app.validator, p.app.config, p.userModule.UserService())
	return nil
}

func (p *moduleProvider) ChatModule() error {
	p.chatModule = chat_module.NewChatModule()
	return nil
}


func (p *moduleProvider) BotModule() error {
	botModule, err := bot_model.NewBotModule(p.app.config, p.app.logger)
	if err != nil {
		return err
	}
	p.botModule = botModule

	// Запуск бота
	go func() {
		if err := p.botModule.Start(); err != nil {
			p.app.logger.Error("Failed to start bot: " + err.Error())
		}
	}()

	return nil
}


