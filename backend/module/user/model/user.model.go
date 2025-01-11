package user_model

import (
	"fmt"

	serv_model "github.com/root9464/Ton-students/module/service_module/model"
)

type SelectedName string
type Role string

const (
	AdminRole   Role = "administarator"
	UserRole    Role = "user"
	CreatorRole Role = "creator"
	ModerRole   Role = "moderator"
)

const (
	Firstname SelectedName = "firstname"
	Lastname  SelectedName = "lastname"
	Nickname  SelectedName = "nickname"
	Username  SelectedName = "username"
)

type User struct {
	ID           int64        `gorm:"primaryKey" json:"id"`
	Username     string       `gorm:"unique;not null" json:"username"`
	Firstname    *string      `json:"firstname"`
	Lastname     *string      `json:"lastname"`
	Nickname     *string      `json:"nickname"`
	SelectedName SelectedName `gorm:"column:selected_name;type:selected_name;not null;default:username" json:"selectedName"`
	Role         Role         `gorm:"column:role;type:role;not null" json:"role"`
	IsPremium    bool         `gorm:"default:false" json:"isPremium"`
	Hash         string       `gorm:"not null" json:"hash"`

	Infos    *[]UserInfo          `gorm:"foreignKey:UserID" json:"infos"`
	Services []serv_model.Service `gorm:"foreignKey:UserId" json:"services"`
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
		"user":    UserRole,
		"moder":   ModerRole,
		"creator": CreatorRole,
		"admin":   AdminRole,
	}

	if role, exists := roles[roleStr]; exists {
		return &role, nil
	}

	return nil, fmt.Errorf("invalid role: %s", roleStr)
}
