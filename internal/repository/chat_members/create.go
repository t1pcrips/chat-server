package chat_members

import (
	"chat-server/internal/client/database"
	"chat-server/internal/model"
	"chat-server/internal/repository/chat_members/converter"
	"chat-server/pkg/errs"
	"context"
	"github.com/Masterminds/squirrel"
)

func (repo *MembersRepositoryImpl) CreateMember(ctx context.Context, info *model.CreateMemberRequest) error {
	repoInfo := converter.ToMemberFromMemberService(info)

	for _, username := range repoInfo.Usernames {
		builderCreateMember := squirrel.Insert(tableMembers).
			PlaceholderFormat(squirrel.Dollar).
			Columns(chatIdColumn, usernameColumn).
			Values(repoInfo.ChatId, username)

		query, args, err := builderCreateMember.ToSql()
		if err != nil {
			return errs.ErrCreateQuery
		}

		q := database.Query{
			Name:     "chat_member repository - create",
			QueryRow: query,
		}

		_, err = repo.db.DB().ExecContext(ctx, q, args...)
		if err != nil {
			return errs.ErrExec
		}

	}

	return nil
}
