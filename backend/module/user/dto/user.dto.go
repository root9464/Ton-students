package user_dto

import user_model "github.com/root9464/Ton-students/module/user/model"

type CreateUserDto struct {
	InitDataRaw string `json:"init-data-raw" validate:"required"`
}

type UpdateUserDto struct {
	NickName     string                  `json:"nick-name"`
	SelectedName user_model.SelectedName `json:"selected-name" validate:"selected_name"`
}

type SrcUser struct {
	AddedToAttachmentMenu bool                   `json:"added_to_attachment_menu"`
	AllowsWriteToPm       bool                   `json:"allows_write_to_pm"`
	FirstName             string                 `json:"first_name"`
	ID                    int64                  `json:"id"`
	IsBot                 bool                   `json:"is_bot"`
	IsPremium             bool                   `json:"is_premium"`
	LastName              string                 `json:"last_name"`
	UserName              string                 `json:"username"`
	LanguageCode          string                 `json:"language_code"`
	PhotoURL              string                 `json:"photo_url"`
	Hash                  string                 `json:"hash"`
	Info                  map[string]interface{} `json:"info"`
}
