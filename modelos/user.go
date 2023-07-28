package modelos

import (
	"time"
)

type User struct {
	Id        int
	Nombre    string
	CreatedAt time.Time
	Status    bool
}

// Referenciamos memoria
func (user *User) AddUser(id int, nombre string, createdAt time.Time, status bool) {
	user.Id = id
	user.Nombre = nombre
	user.CreatedAt = createdAt
	user.Status = status
}
