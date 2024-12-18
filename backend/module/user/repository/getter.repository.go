package user_repository

import (
	"context"

	"github.com/root9464/Ton-students/ent"
	"github.com/root9464/Ton-students/ent/user"
)

func (r *userRepository) GetByID(ctx context.Context, id int64) (*ent.User, error) {
	r.logger.Info("Getting user...")
	user, err := r.db.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		r.logger.Errorf("Error getting user: %v", err)
		return nil, err
	}
	r.logger.Info("User get successfully")
	return user, nil
}
