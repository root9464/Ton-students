package service_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/root9464/Ton-students/ent"
)

func (r *serviceRepository) Create(ctx context.Context, service *ent.Service) (*ent.Service, error) {
	r.logger.Info("Creating service...")

	getService, err := r.db.Service.Create().
		SetUserID(service.UserID).
		SetTitle(service.Title).
		SetDescription(service.Description).
		SetPrice(service.Price).
		Save(ctx)

	if err != nil {
		r.logger.Errorf("Error creating service: %v", err)
		return nil, err
	}
	r.logger.Info("Service create successfully")
	return getService, nil
}

func (r *serviceRepository) CreateTags(ctx context.Context, tagNames []string) ([]*ent.Tags, error) {
	r.logger.Printf("Creating tags: %v", tagNames)

	var createdTags []*ent.Tags

	for _, tagName := range tagNames {
		tag, err := r.db.Tags.Create().
			SetTagName(tagName).
			Save(ctx)

		if err != nil {
			r.logger.Printf("Error creating tag %s: %v", tagName, err)
			return nil, errors.New("Error creating tag")
		}

		r.logger.Printf("Successfully created tag with ID: %v", tag.ID)

		createdTags = append(createdTags, tag)
	}

	if len(createdTags) == 0 {
		return nil, errors.New("No tags created")
	}

	return createdTags, nil
}

func (r *serviceRepository) CreateTag(ctx context.Context, tagName string) (*ent.Tags, error) {
	r.logger.Printf("Creating tag: %s", tagName)

	tag, err := r.db.Tags.Create().
		SetTagName(tagName).
		Save(ctx)

	if err != nil {
		r.logger.Printf("Error creating tag %s: %v", tagName, err)
		return nil, fmt.Errorf("failed to create tag: %w", err)
	}

	r.logger.Printf("Successfully created tag with ID: %v", tag.ID)

	return tag, nil
}
