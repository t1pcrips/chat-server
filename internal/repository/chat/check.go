package chat

import (
	"chat-server/internal/client/database"
	"chat-server/pkg/errs"
	"context"
	"github.com/Masterminds/squirrel"
)

func (repo *ChatRepositoryImpl) CheckChatExists(ctx context.Context, chatId int64) (bool, error) {
	buidlerCheckChat := squirrel.Select("1").
		From(tableChats).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: chatId})

	query, args, err := buidlerCheckChat.ToSql()
	if err != nil {
		return false, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "check chat repository - check chat",
		QueryRow: query,
	}

	var exists string

	err = repo.db.DB().ScanOneContext(ctx, &exists, q, args...)
	if err != nil {
		return false, errs.ErrQueryRowScan
	}

	return exists == "1", nil
}
