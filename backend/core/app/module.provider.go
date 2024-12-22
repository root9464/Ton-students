package app

import (
	auth_module "github.com/root9464/Ton-students/module/auth"
	service_module "github.com/root9464/Ton-students/module/service"
	user_module "github.com/root9464/Ton-students/module/user"
)

type moduleProvider struct {
	authModule    *auth_module.AuthModule
	userModule    *user_module.UserModule
	serviceModule *service_module.ServiceModule

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
		p.ServiceModule,
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
	p.userModule = user_module.NewUserModule(p.app.logger, p.app.validator, p.app.db)
	return nil
}

func (p *moduleProvider) AuthModule() error {
	p.authModule = auth_module.NewAuthModule(p.app.logger, p.app.validator, p.app.config, p.userModule.UserService())
	return nil
}

func (p *moduleProvider) ServiceModule() error {
	p.serviceModule = service_module.NewServiceModule(p.app.logger, p.app.validator, p.app.db)
	return nil
}
