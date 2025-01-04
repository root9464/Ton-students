package user_dto

import user_model "github.com/root9464/Ton-students/module/user/model"

type CreateUserDto struct {
	InitDataRaw string `json:"init-data-raw" validate:"required"`
}

type UserType struct {
	ID           int64                    `json:"id" validate:"required"`
	Username     string                   `json:"username" validate:"required,min=3,max=50"`
	Firstname    *string                  `json:"firstname" `
	Lastname     *string                  `json:"lastname" `
	Nickname     *string                  `json:"nickname" `
	SelectedName *user_model.SelectedName `json:"selectedName"`
	Role         user_model.Role          `json:"role" validate:"required,oneof=administarator user creator moderator"`
	Info         []user_model.UserInfo    `json:"info"`
	IsPremium    bool                     `json:"isPremium"`
	Hash         string                   `json:"hash" validate:"required"`
}

type UserInfoType struct {
	UserId  int64  `json:"userId" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type SelectVisibleNameType struct {
	ID           int64                   `json:"id" validate:"required"`
	SelectedName user_model.SelectedName `json:"selectedName" validate:"required,oneof=firstname lastname nickname username"`
	Hash         string                  `json:"hash" validate:"required"`
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
