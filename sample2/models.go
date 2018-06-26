package app

import (
	"time"
)

//User struct to capture user model
type User struct {
	ID   string `json:"id" datastore:"-"`
	Name string `json:"name"`
}

//Post struct to capture Post model
type Post struct {
	ID        string    `json:"id" datastore:"-"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}
