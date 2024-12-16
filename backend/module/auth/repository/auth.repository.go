package auth_repository

import (
	"context"

	"github.com/root9464/Ton-students/backend/ent"
	"github.com/root9464/Ton-students/backend/ent/user"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func CreateUser(client *ent.Client, user *tma.InitData) (*ent.User, error) {

	newUser, err := client.User.Create().
		SetID(user.User.ID).
		SetFirstName(user.User.FirstName).
		SetLastName(user.User.LastName).
		SetUsername(user.User.Username).
		SetIsPremium(user.User.IsPremium).
		SetHash(user.Hash).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func GetUserByID(client *ent.Client, id int64) (*ent.User, error) {
	user, err := client.User.
		Query().
		Where(user.ID(id)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}
	return user, nil
}
