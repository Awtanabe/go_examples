package models

import "time"


type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Todos     []Todo    `json:"todos" gorm:"foreignKey:UserId;references:ID"`
}

// type User struct {
// 	ID uint `json: "id" gorm: "primaryKey"`
// 	Email string `json: "email" gorm: "unisuq"`
// 	Password string `json: "password"`
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
// 	// Todos []*Todo `gorm: many2many: todos`
// 	// Todos     []*Todo   `gorm:"many2many:user_todos;"`
// 	//  Todos     []Todo    `gorm:"foreignKey:UserID"`
// 	// Todos     []Todo    `json:"todos" gorm:"foreignKey:UserId"`
// 	Todos []Todo `json:"todos" gorm:"foreignKey:UserId;references:ID"`

// }

// リクエストの際の型はないのだろうか
type UserResponse struct {
	 ID uint `json: "id" gorm: "primaryKey"`
	 Email string `json: "email" gorm: "unique"`
}

type UserResponse2 struct {
	ID uint `json: "id" gorm: "primaryKey"`
	Email string `json: "email" gorm: "unique"`
	Todos []Todo `json:"todos"`
}