package app

import auth_module "github.com/root9464/Ton-students/module/auth"

type moduleProvider struct {
	authModule *auth_module.AuthModule

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
		p.AuthModule,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *moduleProvider) AuthModule() error {
	p.authModule = auth_module.NewAuthModule(p.app.logger, p.app.validator, p.app.config)
	return nil
}
