package entity

import "time"

type Role struct {
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement;"`
	Name    string `gorm:"column:name;type:varchar(255)"`
	Deleted bool   `gorm:"column:deleted;not null;default:true;"`

	User []User `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoUpdateTime;type:timestamptz"`
}
