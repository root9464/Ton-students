package serv_model

import (
	"fmt"

	"gorm.io/gorm"
)

type Service struct {
	ID     string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId int64   `gorm:"not null;" json:"userId"`
	Price  float64 `gorm:"not null" json:"price"`

	Infos    []ServiceInfo   `gorm:"foreignKey:ServiceId" json:"infos"`
	Tags     *[]Tags         `gorm:"foreignKey:ServiceId" json:"tags"`
	Settings ServiceSettings `gorm:"foreignKey:ServiceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"settings"`
}

func (s *Service) AfterCreate(tx *gorm.DB) (err error) {
	settingsInDb := new(ServiceSettings)
	if err := tx.Where("service_id = ?", s.ID).First(settingsInDb).Error; err == nil {
		return nil
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking existing ServiceSettings for service ID %s: %v", s.ID, err)
	}

	defaultSettings := ServiceSettings{
		ServiceId:          s.ID,
		ColorHeader:        "#007aff",
		IsPrepayment:       false,
		IsDisabled:         false,
		IsAdditionalButton: false,
	}

	if err := tx.Create(&defaultSettings).Error; err != nil {
		return fmt.Errorf("failed to create default service settings for service ID %s: %v", s.ID, err)
	}
	return nil
}

type ServiceInfo struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	ServiceId string `gorm:"type:uuid;not null;" json:"serviceId"`
	Title     string `gorm:"not null" json:"title"`
	Content   string `gorm:"type:text" json:"content"`
}

type Tags struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	ServiceId string `gorm:"type:uuid;not null;" json:"serviceId"`
	Name      string `gorm:"not null" json:"name"`
}

type ServiceSettings struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	ServiceId string `gorm:"type:uuid;not null;unique;index" json:"serviceId"`

	ColorHeader string  `gorm:"not null;default:#007aff" json:"colorHeader"`
	ButtonText  *string `gorm:"null" json:"buttonText"`

	IsPrepayment       bool `gorm:"not null; default:false" json:"isPrepayment"`
	IsDisabled         bool `gorm:"not null; default:false" json:"isDisabled"`
	IsAdditionalButton bool `gorm:"not null; default:false" json:"isAdditionalButton"`
}
