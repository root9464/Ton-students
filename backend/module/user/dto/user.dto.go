package user_dto

import user_model "github.com/root9464/Ton-students/module/user/model"

type CreateUserDto struct {
	InitDataRaw string `json:"init-data-raw" validate:"required"`
}

type UserType struct {
	ID           int64                   `json:"id" validate:"required"`
	Username     string                  `json:"username" validate:"required,min=3,max=50"`
	Firstname    *string                 `json:"firstname" `
	Lastname     *string                 `json:"lastname" `
	Nickname     *string                 `json:"nickname" `
	SelectedName user_model.SelectedName `json:"selectedName" validate:"required,oneof=firstname lastname nickname username"`
	Role         user_model.Role         `json:"role" validate:"required,oneof=administarator user creator moderator"`
	Info         []user_model.UserInfo   `json:"info"`
	IsPremium    bool                    `json:"isPremium"`
	Hash         string                  `json:"hash" validate:"required"`
}
