package chat_members

import (
	"github.com/t1pcrips/chat-service/internal/repository"
	"github.com/t1pcrips/platform-pkg/pkg/database"
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
