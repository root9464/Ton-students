package user_dto

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserDto struct {
	InitDataRaw string `json:"init-data-raw" validate:"required"`
}

type ShortUserType struct {
	ID          int64          `json:"id" validate:"required"`
	VisibleName string         `json:"username" validate:"required"`
	Role        string         `json:"role" validate:"required"`
	Hash        string         `json:"hash" validate:"required"`
	Infos       []UserInfoType `json:"infos" validate:"required"`
}

type UserType struct {
	ID           int64          `json:"id" validate:"required"`
	Username     string         `json:"username" validate:"required,min=3,max=50"`
	Firstname    *string        `json:"firstname" `
	Lastname     *string        `json:"lastname" `
	Nickname     *string        `json:"nickname" `
	SelectedName *string        `json:"selectedName"`
	Info         []UserInfoType `json:"info"`
	IsPremium    bool           `json:"isPremium"`
	Hash         string         `json:"hash" validate:"required"`
}

type UserInfoType struct {
	UserId  int64  `json:"userId" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type SelectVisibleNameType struct {
	ID           int64  `json:"id" validate:"required"`
	SelectedName string `json:"selectedName" validate:"required,oneof=firstname lastname nickname username"`
	Hash         string `json:"hash" validate:"required"`
}

type SetUserNicknameType struct {
	ID       int64  `json:"id" validate:"required"`
	Nickname string `json:"nickname" validate:"required,min=3,max=50"`
	Hash     string `json:"hash" validate:"required"`
}

type UpdateUserInfoType struct {
	ID      string `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Hash    string `json:"hash" validate:"required"`
}

type DeleteUserInfoType struct {
	ID   string `json:"id" validate:"required"`
	Hash string `json:"hash" validate:"required"`
}

func MaxInfos(fl validator.FieldLevel) bool {
	if field, ok := fl.Field().Interface().([]UserInfosType); ok {
		return len(field) <= 3
	}
	return false
}

type ManyUserInfoType struct {
	UserId int64           `json:"userId" validate:"required"`
	Infos  []UserInfosType `json:"infos" validate:"required,maxinfos"`
}

type UserInfosType struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
