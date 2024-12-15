package models

import (
	"time"
)

// type Todo struct {
// 	ID uint `json: "id" gorm: "primaryKey"`
// 	Title string `json: "title" gorm: "not null"`
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
// 	User      User      `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
// 	UserId    uint      `json:"user_id" gorm:"not null"`
// }

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

type TodoResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}