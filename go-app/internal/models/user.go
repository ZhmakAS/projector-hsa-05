package models

import "time"

type User struct {
	Id        int64     `json:"id" bson:"id" db:"id" fake:"{number:1,10}"`
	FirstName string    `json:"first_name" db:"first_name" fake:"{firstname}"`
	LastName  string    `json:"last_name" db:"last_name" fake:"{lastname}"`
	Phone     string    `json:"phone" db:"phone" fake:"{phone}"`
	CreatedAt time.Time `json:"created_at" db:"created_at" fake:"{year}-{month}-{day}"`
}
