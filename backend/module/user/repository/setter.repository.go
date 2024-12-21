package user_repository

import (
	"context"
	"fmt"

	"github.com/root9464/Ton-students/ent"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func (r *userRepository) Create(ctx context.Context, user *tma.InitData) (*ent.User, error) {
	r.logger.Info("Creating user...")
	getUser, err := r.db.User.Create().
		SetID(user.User.ID).
		SetFirstName(user.User.FirstName).
		SetLastName(user.User.LastName).
		SetUserName(user.User.Username).
		SetIsPremium(user.User.IsPremium).
		SetHash(user.Hash).
		Save(ctx)
	if err != nil {
		r.logger.Errorf("Error creating user: %v", err)
		return nil, err
	}
	r.logger.Info("User create successfully")
	return getUser, nil
}

func (r *userRepository) Update(ctx context.Context, user *tma.InitData) (*ent.User, error) {
	getUser, err := r.db.User.UpdateOneID(user.User.ID).
		SetFirstName(user.User.FirstName).
		SetLastName(user.User.LastName).
		SetUserName(user.User.Username).
		SetIsPremium(user.User.IsPremium).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, fmt.Errorf("user not found")
	case err != nil:
		return nil, err
	}

	return getUser, nil
}
