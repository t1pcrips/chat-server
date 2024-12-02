package message_model

import "time"

type CreateMessageRequest struct {
	ChatId    int64
	From      string
	Text      string
	Timestamp time.Time
}
