package users

import "time"

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name" binding:"required,min=3,max=255"`
	Email     string    `json:"email" bson:"email" binding:"required,email,min=3,max=255"`
	Password  string    `json:"password" bson:"password,omitempty" binding:"required,min=6,max=255"`
	Token     string    `json:"token" bson:"token,omitempty" binding:"omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty" binding:"omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty" binding:"omitempty"`
}
