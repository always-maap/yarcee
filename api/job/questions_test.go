package main

import (
	"api/models"
	"testing"
)

func TestParseQuestion(t *testing.T) {
	content := `"""
	1. Two Sum
Easy
Array | Hash Table
---
Given Question
Solve it!
"""
class Solution:
	def twoSum():
		pass
`

	got := parseQuestion(content)
	want := models.Question{
		No:         1,
		Name:       "Two Sum",
		Difficulty: "Easy",
		Subject:    "Array | Hash Table",
		Solution: `class Solution:
	def twoSum():
		pass`,
		Problem: `Given Question
Solve it!`,
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
