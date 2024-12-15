package controller

import "github.com/root9464/Ton-students/backend/module/auth/service"

type Controller struct {
	authService service.IAuthService
}

func NewController(authService service.IAuthService) *Controller {
	return &Controller{
		authService: authService,
	}
}
