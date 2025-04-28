package dao

import "time"

type AuthRole string

var (
	Admin AuthRole = "admin"
	User  AuthRole = "user"
)

type AuthEntity struct {
	ID           uint     `gorm:"primaryKey"`
	Name         string   `gorm:"type:varchar(255);not null"`
	Email        string   `gorm:"unique;not null"`
	Password     string   `gorm:"type:text;not null"`
	Avatar       string   `gorm:"default:'http://localhost:3000/public/propil.png'"`
	Role         AuthRole `gorm:"not null;type:enum('admin','user');default:'user';"`
	RefreshToken *string  `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Store         []StoreEntity `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (AuthEntity) TableName() string {
	return "auth"
}
