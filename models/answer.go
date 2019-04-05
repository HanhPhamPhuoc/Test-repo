package models

import "time"

// Answer
type Answer struct {
	Id        int
	Content   string
	Owner     string
	CreatedAt time.Time
}
