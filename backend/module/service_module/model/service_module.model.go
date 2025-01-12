package serv_model

type Service struct {
	ID     string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId int64   `gorm:"not null;" json:"userId"`
	Price  float64 `gorm:"not null" json:"price"`

	Infos    []ServiceInfo   `gorm:"foreignKey:ServiceId" json:"infos"`
	Tags     *[]Tags         `gorm:"foreignKey:ServiceId" json:"tags"`
	Settings ServiceSettings `gorm:"foreignKey:ServiceId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"settings"`
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
