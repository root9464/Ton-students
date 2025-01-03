package auth_dto

import user_model "github.com/root9464/Ton-students/module/user/model"

type AutorizeDto struct {
	InitDataRaw string `json:"init-data-raw" validate:"required"`
}

type UserType struct {
	Username     string                  `json:"username" validate:"required,min=3,max=50"`
	Firstname    string                  `json:"firstname" validate:"required,max=50"`
	Lastname     string                  `json:"lastname" validate:"required,max=50"`
	Nickname     *string                 `json:"nickname" validate:"omitempty,max=50"`
	SelectedName user_model.SelectedName `json:"selectedName" validate:"required,oneof=firstname lastname nickname username"`
	Role         user_model.Role         `json:"role" validate:"required,oneof=administarator user creator moderator"`
	Info         []user_model.UserInfo   `json:"info"`
	IsPremium    bool                    `json:"isPremium"`
	Hash         string                  `json:"hash" validate:"required"`
}
