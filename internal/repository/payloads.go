package repository

import "time"

type SendMessageRequest struct {
	From     string
	Text     string
	TimeSend time.Time
}
