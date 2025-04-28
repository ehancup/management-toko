package dao

import "time"

type ProductEntity struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(255)"`
	Price uint
	Image string `gorm:"default:'http://localhost:3000/public/propil.png'"`
	StoreID uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProductEntity) TableName() string {
	return "product"
}
