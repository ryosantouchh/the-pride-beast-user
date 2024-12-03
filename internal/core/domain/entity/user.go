package entity

import "time"

// type Role string

// const (
// 	SuperAdminRole Role = "super-admin"
// 	AdminRole      Role = "admin"
// 	MemberRole     Role = "member"
// 	InstructorRole Role = "instructor"
// )

type GenderEnum string

const (
	Male        GenderEnum = "male"
	Female      GenderEnum = "female"
	NotIdentify GenderEnum = "not-identify"
)

type User struct {
	ID          uint       `gorm:"column:id;primaryKey;autoIncrement;"`
	Username    string     `gorm:"column:username;type:varchar(255);unique;not null;"`
	Password    string     `gorm:"column:password;type:varchar(255);not null;"`
	Firstname   string     `gorm:"column:firstname;type:varchar(255);not null;"`
	Lastname    string     `gorm:"column:lastname;type:varchar(255);not null;"`
	Age         uint       `gorm:"column:age;"`
	Email       string     `gorm:"column:email;type:varchar(255);unique;not null;"`
	PhoneNumber string     `gorm:"column:phonenumber;type:varchar(32);unique;"`
	Gender      GenderEnum `gorm:"column:gender;type:varchar(255);"`
	RoleID      uint       `gorm:"column:role_id;"`

	Deleted bool `gorm:"column:deleted;not null;default:true;"`

	OauthProvider string `gorm:"column:oauth_provider;type:varchar(255);"`
	OauthUID      string `gorm:"column:oauth_uid;type:varchar(255);"`
	RefreshToken  string `gorm:"column:refresh_token;type:varchar(255);"`

	// Bookings      []Booking     `gorm:"foreignKey:UserID;"`
	// UserCredits   []UserCredit  `gorm:"foreignKey:UserID;"`
	// CreditUsed    []CreditUsed  `gorm:"foreignKey:UserID;"`
	// Orders        []Order       `gorm:"foreignKey:UserID;"`
	// UserAddresses []UserAddress `gorm:"foreignKey:UserID;"`

	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoUpdateTime;type:timestamptz"`
}
