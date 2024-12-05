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
	Usernames []string
}

type CreateChatRequest struct {
	Usernames []string
}

type User struct {
	Username string
}
