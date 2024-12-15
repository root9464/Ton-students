package auth_repository

import (
	"context"
	ent "root/database"
)

type Repository struct {
	client *ent.Client
}

type RepositoryInterface interface {
	CreateUser(user *ent.User) (*ent.User, error)
	GetUser(username string) (*ent.User, error)
	UpdateUser(user *ent.User) (*ent.User, error)
	DeleteUser(user *ent.User) error
}

func NewRepository(client *ent.Client) RepositoryInterface {
	return &Repository{
		client: client,
	}
}

func (rep *Repository) CreateUser(user *ent.User) (*ent.User, error) {
	newUser, err := rep.client.User.Create().
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetUsername(user.Username).
		SetIsPremium(user.IsPremium).
		SetRole(user.Role).
		SetHash(user.Hash).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (rep *Repository) DeleteUser(user *ent.User) error {
	panic("unimplemented")
}

func (rep *Repository) GetUser(username string) (*ent.User, error) {
	panic("unimplemented")
}

func (rep *Repository) UpdateUser(user *ent.User) (*ent.User, error) {
	panic("unimplemented")
}
