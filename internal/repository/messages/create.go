package messages

import (
	"chat-server/internal/client/database"
	"chat-server/internal/model"
	"chat-server/internal/repository/messages/converter"
	"chat-server/pkg/errs"
	"context"
	"github.com/Masterminds/squirrel"
)

func (repo *MessagesRepositoryImpl) CreateMessage(ctx context.Context, info *model.CreateMessageRequest) error {
	repoInfo := converter.ToMessageFromMessageService(info)
	builderCreateMessage := squirrel.Insert(tableMessages).
		PlaceholderFormat(squirrel.Dollar).
		Columns(chatIdColumn, fromColumn, textColumn, timestampColumn).
		Values(repoInfo.ChatId, repoInfo.From, repoInfo.Text, repoInfo.Timestamp)

	query, args, err := builderCreateMessage.ToSql()
	if err != nil {
		return errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "message repository - createMessage",
		QueryRow: query,
	}

	result, err := repo.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errs.ErrExec
	}

	if result.RowsAffected() == 0 {
		return errs.ErrNoRows
	}

	return nil
}
