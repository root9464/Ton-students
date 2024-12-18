package user_repository

import (
	"context"
	"fmt"

	"github.com/root9464/Ton-students/ent"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func (r *userRepository) Create(ctx context.Context, user *tma.InitData) error {
	r.logger.Info("Creating user...")
	_, err := r.db.User.Create().
		SetID(user.User.ID).
		SetFirstName(user.User.FirstName).
		SetLastName(user.User.LastName).
		SetUsername(user.User.Username).
		SetIsPremium(user.User.IsPremium).
		SetHash(user.Hash).
		Save(ctx)
	if err != nil {
		r.logger.Errorf("Error creating user: %v", err)
		return err
	}
	r.logger.Info("User create successfully")
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *tma.InitData) error {
	err := r.db.User.UpdateOneID(user.User.ID).
		SetFirstName(user.User.FirstName).
		SetLastName(user.User.LastName).
		SetUsername(user.User.Username).
		SetIsPremium(user.User.IsPremium).
		Exec(ctx)

	switch {
	case ent.IsNotFound(err):
		return fmt.Errorf("user not found")
	case err != nil:
		return err
	}

	return nil
}
