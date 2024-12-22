package custom_validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/root9464/Ton-students/ent/user"
)

func IsValidSelectedName(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	switch user.SelectedName(value) {
	case user.SelectedNameFirstname, user.SelectedNameLastname, user.SelectedNameNickname, user.SelectedNameUsername:
		return true
	default:
		return false
	}
}
