package database

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		&model.User{},
		&model.Role{},
		&model.Region{},
		&model.Category{},
		&model.Characteristic{},
		&model.Product{},
		&model.CharacteristicValue{},
		&model.Image{},
		&model.Price{},
		&model.Order{},
		&model.PaymentMethod{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	if err := createDefaultRole(db); err != nil {
		return err
	}

	if err := createRegions(db); err != nil {
		return err
	}

	return nil
}

func createDefaultRole(db *gorm.DB) error {
	roles := []model.Role{
		{Name: "admin"},
		{Name: "user"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, model.Role{Name: role.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}

func createRegions(db *gorm.DB) error {
	var existingRegions []model.Region

	if err := db.Find(&existingRegions).Error; err != nil {
		return err
	}

	existingRegionMap := make(map[string]struct{})
	for _, region := range existingRegions {
		existingRegionMap[region.Name] = struct{}{}
	}

	var newRegions []model.Region
	for _, region := range constants.Regions {
		if _, exists := existingRegionMap[region]; !exists {
			newRegions = append(newRegions, model.Region{Name: region})
		}
	}

	if len(newRegions) > 0 {
		if err := db.Create(&newRegions).Error; err != nil {
			return err
		}
	}

	return nil
}
