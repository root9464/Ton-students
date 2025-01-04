package user_funcs

import (
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/samber/lo"
)

func GetVisibleName(newUser *user_model.User) string {
	nameMap := lo.Assign(
		map[user_model.SelectedName]string{
			user_model.Username: newUser.Username,
		},
		lo.OmitBy(map[user_model.SelectedName]string{
			user_model.Firstname: lo.FromPtr(newUser.Firstname),
			user_model.Lastname:  lo.FromPtr(newUser.Lastname),
			user_model.Nickname:  lo.FromPtr(newUser.Nickname),
		}, func(_ user_model.SelectedName, value string) bool {
			return value == ""
		}),
	)

	entry, found := lo.Find(lo.Entries(nameMap), func(entry lo.Entry[user_model.SelectedName, string]) bool {
		return entry.Key == newUser.SelectedName
	})

	if found {
		return entry.Value
	}

	return "none"
}
