package user_model

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
	Firstname    string       `json:"firstname"`
	Lastname     string       `json:"lastname"`
	Nickname     string       `json:"nickname"`
	SelectedName SelectedName `gorm:"column:selected_name;type:selected_name" json:"selectedName"`
	Role         Role         `gorm:"column:role;type:role;not null" json:"role"`
	Info         string       `json:"info"`
	IsPremium    bool         `gorm:"default:false" json:"isPremium"`
	Hash         string       `gorm:"not null" json:"hash"`
}
