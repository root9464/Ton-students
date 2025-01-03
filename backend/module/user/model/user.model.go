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
	Nickname     *string      `json:"nickname"`
	SelectedName SelectedName `gorm:"column:selected_name;type:selected_name" json:"selectedName"`
	Role         Role         `gorm:"column:role;type:role;not null" json:"role"`
	Infos        *[]UserInfo  `gorm:"foreignKey:UserID" json:"infos"`
	IsPremium    bool         `gorm:"default:false" json:"isPremium"`
	Hash         string       `gorm:"not null" json:"hash"`
}

type UserInfo struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	UserID  int64  `gorm:"not null;index" json:"userId"`
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
