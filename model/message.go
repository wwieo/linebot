package model

import "time"

type Message struct {
	MessageID string     `form:"messageID";bson:"messageid"`
	UserID    string     `form:"userID";bson:"userID"`
	Username  string     `form:"username";bson:"username"`
	Text      string     `form:"text";bson:"text"`
	Timestamp *time.Time `form:"timestamp";bson:"timestamp"`
}
