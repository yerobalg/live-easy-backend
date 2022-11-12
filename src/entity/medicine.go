package entity

import (
	"gorm.io/gorm"
)

type Medicine struct {
	// Basic Fields
	ID        int64          `gorm:"primary_key" json:"id"`
	CreatedAt int64          `json:"createdAt"`
	UpdatedAt int64          `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy *int64         `json:"createdBy" gorm:"default:null"`
	UpdatedBy *int64         `json:"updatedBy" gorm:"default:null"`
	DeletedBy *int64         `json:"deletedBy" gorm:"default:null"`

	Name     string `json:"name" gorm:"not null;type:varchar(255)"`
	Price    int64  `json:"price" gorm:"not null;"`
	Quantity int64  `json:"quantity" gorm:"not null;"`
	ImageURL string `json:"imageURL" gorm:"not null;type:varchar(255)"`
	UserID   int64  `json:"userID" gorm:"index;not null"`
}

type MedicineParam struct {
	ID     int64 `uri:"id" param:"id"`
	UserID int64 `json:"-" param:"userID"`
	PaginationParam
}

type MedicineInputParam struct {
	Name     string `json:"name" binding:"required"`
	Price    int64  `json:"price" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
}

type MedicineUpdateInputParam struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
}
