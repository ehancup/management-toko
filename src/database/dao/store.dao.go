package dao

import (
	"time"

	"gorm.io/gorm"

	"gin-boilerplate/src/utils/logger"
)

// StoreEntity represents the store table in the database
type StoreEntity struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(255)"`
	Description *string `gorm:"type:text"`
	Avatar      string  `gorm:"default:'http://localhost:3000/public/propil.png'"`
	UserID      uint

	CreatedAt time.Time
	UpdatedAt time.Time
	Product   []ProductEntity `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName sets the insert table name for this struct type
func (StoreEntity) TableName() string {
	return "store"
}

// BeforeCreate is a GORM hook that is called before a new record is inserted into the database
func (StoreEntity) BeforeCreate(tx *gorm.DB) (err error) {
	logger.Info("saving toko")
	return nil
}
