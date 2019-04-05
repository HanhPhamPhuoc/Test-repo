package models

import "time"

// Question
type Question struct {
	Id        int
	Content   string
	Owner     string
	CreatedAt *time.Time
}
