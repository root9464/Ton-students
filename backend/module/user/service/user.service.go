package user_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_model "github.com/root9464/Ton-students/module/user/model"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IUserService = (*userService)(nil)

type ResponseUserData struct {
	ID           int64                   `json:"id"`
	Visiblename  string                  `json:"visiblename"`
	SelectedName user_model.SelectedName `json:"selectedName"`
	Role         user_model.Role         `json:"role"`
	Infos        *[]user_model.UserInfo  `json:"infos"`
	IsPremium    bool                    `json:"isPremium"`
	Hash         string                  `json:"hash"`
}

type ResponseCreateUser struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    *ResponseUserData `json:"data"`
}

type IUserService interface {
	Create(ctx context.Context, dto *user_dto.UserType) (*ResponseCreateUser, error)
	SelectVisibleName(ctx context.Context, dto *user_dto.SelectVisibleNameType) error

	AddUserInfo(ctx context.Context, dto *user_dto.UserInfoType) error
}

type userService struct {
	logger    *logger.Logger
	validator *validator.Validate

	repo user_repository.IUserRepository
}

func NewUserService(
	logger *logger.Logger,
	validator *validator.Validate,
	repo user_repository.IUserRepository,
) *userService {
	return &userService{
		logger:    logger,
		validator: validator,
		repo:      repo,
	}
}
