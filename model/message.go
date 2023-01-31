package model

import "time"

type Message struct {
	MessageID string    `bson:"messageid"`
	UserID    string    `bson:"userid"`
	UserName  string    `bson:"username"`
	Text      string    `bson:"text"`
	Timestamp time.Time `bson:"timestamp"`
}
