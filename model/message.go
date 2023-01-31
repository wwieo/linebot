package model

import "time"

type Message struct {
	MessageID string
	UserID    string
	UserName  string
	Text      string
	Timestamp time.Time
}
