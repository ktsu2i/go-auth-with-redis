package model

import "time"

type User struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
