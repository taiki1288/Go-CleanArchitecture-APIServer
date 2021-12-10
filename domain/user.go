package domain

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
}

type Users []User