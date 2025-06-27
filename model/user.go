package model

type User struct {
	ID             string `bson:"_id"`
	Username       string `bson:"username"`
	HashedPassword string `bson:"hashed_password"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
