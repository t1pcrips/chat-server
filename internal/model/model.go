package model

import "time"

type CreateMessageRequest struct {
	ChatId    int64
	From      string
	Text      string
	Timestamp time.Time
}

type CreateMemberRequest struct {
	ChatId    int64
	Usernames []User
}

type User struct {
	Username string
}
