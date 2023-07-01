package model

import (
	"time"
)

type Sandbox struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	Language     string    `json:"language"`
	Code         string    `json:"code"`
	Status       string    `json:"status"`
	StdOut       string    `json:"stdout"`
	StdErr       string    `json:"stderr"`
	ExecDuration int       `json:"execDuration" gorm:"type:integer; null"`
	ExecMemUse   int       `json:"execMemUse" gorm:"type:integer; null"`
	UserRefer    uint      `json:"userRefer"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
