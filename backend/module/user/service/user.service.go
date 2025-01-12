package user_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_repository "github.com/root9464/Ton-students/module/user/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IUserService = (*userService)(nil)

type IUserService interface {
	UpsertUser(ctx context.Context, dto *user_dto.UserType) (*user_dto.ShortUserType, error)

	SelectVisibleName(ctx context.Context, dto *user_dto.SelectVisibleNameType) error
	SetUserNickname(ctx context.Context, dto *user_dto.SetUserNicknameType) error
	AddUserInfo(ctx context.Context, dto *user_dto.UserInfoType) error
	UpdateUserInfo(ctx context.Context, dto *user_dto.UpdateUserInfoType) error
	DeleteUserInfo(ctx context.Context, dto *user_dto.DeleteUserInfoType) error
	AddManyUserInfo(ctx context.Context, dto *user_dto.ManyUserInfoType) error
}

type userService struct {
	logger    *logger.Logger
	validator *validator.Validate

	repo user_repository.IUserRepository
}

func NewUserService(logger *logger.Logger, validator *validator.Validate, repo user_repository.IUserRepository) *userService {
	return &userService{logger: logger, validator: validator, repo: repo}
}
