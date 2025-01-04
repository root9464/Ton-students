package serv_model

type Service struct {
	ID          string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId      int64   `gorm:"not null;" json:"userId"` // ID пользователя без импорта пакета user_model
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Tags        []Tags  `gorm:"foreignKey:ServiceId" json:"tags"`
}

type Tags struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	ServiceId string `gorm:"type:uuid;not null;" json:"serviceId"`
	Name      string `gorm:"not null" json:"name"`
}
