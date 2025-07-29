package model

import "time"

type User struct {
	ID             string `bson:"_id"`
	Username       string `bson:"username"`
	HashedPassword string `bson:"hashed_password"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreatedUser struct {
	Id             string    `json:"id" bson:"_id"`
	CompanyId      string    `json:"company_id" bson:"company_id"`
	Username       string    `json:"username" bson:"username"`
	HashedPassword string    `json:"hashed_password" bson:"hashed_password"`
	Email          string    `json:"email" bson:"email"`
	FullName       string    `json:"full_name" bson:"full_name"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

type CreateUserRequest struct {
	CompanyId string `json:"company_id" binding:"omitempty,uuid"`
	Username  string `json:"username" binding:"required,min=3,max=30,valid_user_username"`
	Password  string `json:"password" binding:"required,min=8,max=32,valid_user_password"`
	Email     string `json:"email" binding:"required,email"`
	FullName  string `json:"full_name" binding:"omitempty,min=1,max=30,valid_name"`
}
