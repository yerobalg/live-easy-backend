package entity

import (
	"gorm.io/gorm"
)

type User struct {
	// Basic Fields
	ID        int64          `gorm:"primary_key" json:"id"`
	CreatedAt int64          `json:"createdAt"`
	UpdatedAt int64          `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy *int64         `json:"createdBy" gorm:"default:null"`
	UpdatedBy *int64         `json:"updatedBy" gorm:"default:null"`
	DeletedBy *int64         `json:"deletedBy" gorm:"default:null"`

	Email           string `json:"email" gorm:"not null;unique;type:varchar(255)"`
	Password        string `json:"-" gorm:"type:text"`
	Name            string `json:"name" gorm:"not null;type:varchar(255)"`
	IsGoogleAccount bool   `json:"isGoogleAccount" gorm:"not null;default:false"`
}

type UserParam struct {
	ID    int64  `uri:"id" param:"id"`
	Email string `json:"-" param:"email"`
	PaginationParam
}

type UserLoginInputParam struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginGoogleInputParam struct {
	FirebaseJWT string `json:"firebaseJWT" binding:"required"`
}

type UserRegisterInputParam struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UserLoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
