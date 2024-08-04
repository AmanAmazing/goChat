package models

import "time"

type PostLogin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	ID        int
	Title     string
	Text      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
