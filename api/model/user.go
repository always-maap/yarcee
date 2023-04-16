package model

import (
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Username  string    `gorm:"unique" json:"username"`
	Password  []byte    `json:"-"`
	Sandboxes []Sandbox `gorm:"foreignKey:UserRefer" json:"sandboxes"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
