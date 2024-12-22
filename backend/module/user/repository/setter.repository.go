package user_repository

import (
	"context"
	"fmt"

	"github.com/root9464/Ton-students/ent"
)

func (r *userRepository) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	r.logger.Info("Creating user...")
	getUser, err := r.db.User.Create().
		SetID(user.ID).
		SetFirstname(user.Firstname).
		SetLastname(user.Lastname).
		SetUsername(user.Username).
		SetIsPremium(user.IsPremium).
		SetHash(user.Hash).
		SetUsername(user.Username).
		Save(ctx)
	if err != nil {
		r.logger.Errorf("Error creating user: %v", err)
		return nil, err
	}
	r.logger.Info("User create successfully")
	return getUser, nil
}

func (r *userRepository) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	getUser, err := r.db.User.UpdateOneID(user.ID).
		SetFirstname(user.Firstname).
		SetLastname(user.Lastname).
		SetUsername(user.Username).
		SetIsPremium(user.IsPremium).
		SetHash(user.Hash).
		Save(ctx)

	switch {
	case ent.IsNotFound(err):
		return nil, fmt.Errorf("user not found")
	case err != nil:
		return nil, err
	}

	return getUser, nil
}

// func (r *userRepository) UpdateInfo(ctx context.Context, id int64, dto *user_dto.UpdateUserDto) (*ent.User, error) {
// }
