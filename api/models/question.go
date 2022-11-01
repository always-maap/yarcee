package models

type Question struct {
	Id         uint   `json:"id"`
	No         int    `json:"no"`
	Name       string `json:"name"`
	Subject    string `json:"subject"`
	Difficulty string `json:"difficulty"`
	Solution   string `json:"solution"`
	Problem    string `json:"problem"`
}
