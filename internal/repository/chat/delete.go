package chat

import (
	"chat-server/internal/client/database"
	"chat-server/pkg/errs"
	"context"
	"github.com/Masterminds/squirrel"
)

func (repo *ChatRepositoryImpl) DeleteChat(ctx context.Context, chatId int64) error {
	builderDeleteChat := squirrel.Delete(tableChats).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: chatId})

	query, args, err := builderDeleteChat.ToSql()
	if err != nil {
		return errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "chat repository - deleteChat",
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
