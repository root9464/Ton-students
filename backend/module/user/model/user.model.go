package user_model

import (
	"fmt"

	notifications_model "github.com/root9464/Ton-students/module/notifications/model"
)

type SelectedName string
type Role string

const (
	AdminRole   Role = "administarator"
	UserRole    Role = "user"
	CreatorRole Role = "creator"
	ModerRole   Role = "moderator"
)

type User struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"unique;not null" json:"username"`
	Nickname  string `gorm:"unique;not null" json:"nickname"`
	Role      Role   `gorm:"column:role;type:role;not null;default:user" json:"role"`
	IsPremium bool   `gorm:"default:false" json:"isPremium"`
	Hash      string `gorm:"not null" json:"hash"`

	Infos         []UserInfo                         `gorm:"foreignKey:UserID" json:"infos"`
	Notifications []notifications_model.Notification `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"notifications"`
}

type UserInfo struct {
	ID      string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID  int64  `gorm:"not null;index" json:"userId"`
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func ParseRole(roleStr string) (*Role, error) {
	roles := map[string]Role{
		"user":           UserRole,
		"moderator":      ModerRole,
		"creator":        CreatorRole,
		"administarator": AdminRole,
	}

	if role, exists := roles[roleStr]; exists {
		return &role, nil
	}

	return nil, fmt.Errorf("invalid role: %s", roleStr)
}
