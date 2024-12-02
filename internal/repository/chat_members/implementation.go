package chat_members

import (
	"chat-server/internal/client/database"
	"chat-server/internal/repository"
)

const (
	tableMembers   = "chat_members"
	chatIdColumn   = "chat_id"
	usernameColumn = "username"
)

type MembersRepositoryImpl struct {
	db database.Client
}

func NewMembersRepository(db database.Client) repository.MembersRepository {
	return &MembersRepositoryImpl{
		db: db,
	}
}
