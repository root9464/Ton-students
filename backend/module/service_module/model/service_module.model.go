package serv_model

import (
	chat_model "github.com/root9464/Ton-students/module/chat/model"
	user_model "github.com/root9464/Ton-students/module/user/model"
)

type Service struct {
	ID     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID int64  `gorm:"not null;" json:"userID"`

	Price float64       `gorm:"not null" json:"price"`
	Infos []ServiceInfo `gorm:"foreignKey:ServiceId" json:"infos"`
	Tags  []Tags        `gorm:"foreignKey:ServiceId" json:"tags"`

	User     user_model.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	Settings ServiceSettings `gorm:"foreignKey:ServiceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"settings"`

	ChatID *string          `gorm:"type:uuid;uniqueIndex" json:"chatId"` // Уникальный индекс для связи one-to-one
	Chat   *chat_model.Chat `gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"chat"`
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
