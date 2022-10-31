package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `gorm:"unique" json:"username"`
	Password []byte `json:"-"`
}
