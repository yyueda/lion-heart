package models

import "time"

type User struct {
	ID       string
	Name     string
	Email    string
	CreateAt time.Time
}
