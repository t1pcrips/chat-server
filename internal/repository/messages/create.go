package messages

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/repository/messages/converter"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"github.com/t1pcrips/platform-pkg/pkg/database"
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
