package model

import (
	"time"
)

type Sandbox struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Language  string    `json:"language"`
	Code      string    `json:"code"`
	UserRefer uint      `json:"userRefer"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
