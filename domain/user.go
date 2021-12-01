package domain

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Users []User